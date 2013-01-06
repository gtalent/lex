/*
   Copyright 2011-2013 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package lex

import (
	"strconv"
	"strings"
)

//Returns tokens from a generic parser used in another project.
//Here mainly for compatibility reasons.
func Tokens(input string) []Token {
	var tokens []Token

	symbols := []string{"&&", "||", "=<", "=>", "==", "!=", "<", ">", "/", "*", "-", "+", "(", ")", "!"}
	stringTypes := []Pair{{Opener: "'", Closer: "'"}, {Opener: "\"", Closer: "\""}}
	commentTypes := []Pair{{Opener: "#", Closer: "\n"}}
	lex := NewAnalyzer(symbols, []string{}, stringTypes, commentTypes, true)
	for point := 0; point < len(input); {
		var t Token
		t.TokType, t.TokValue, point = lex.NextToken(input, point)
		if t.TokType == IntLiteral {
			t.TokValue, _ = strconv.Atoi(t.TokValue.(string))
		} else if t.TokType == BoolLiteral {
			t.TokValue, _ = strconv.ParseBool(t.TokValue.(string))
		}
		tokens = append(tokens, t)
	}

	return tokens
}

func isCharacter(val byte) bool {
	return (val < 91 && val > 64) || (val < 123 && val > 96)
}

func isNumber(val byte) bool {
	return (47 < val && val < 58)
}

type Pair struct {
	Opener string
	Closer string
}

//Returns the contents of the pair.
func (me *Pair) parse(val string) string {
	token := ""
	for len(val) != 0 && val[:len(me.Closer)] != me.Closer {
		val = val[1:]
		token += string(val[0])
	}
	return token
}

func (me *Pair) opens(val string) bool {
	return len(me.Opener) <= len(val) && val[:len(me.Opener)] == me.Opener
}

func isWhitespace(val byte) bool {
	return val == ' ' || val == '\t' || val == '\n'
}

type LexAnalyzer struct {
	identTable   []string
	numLitTable  []string
	keywords     []string
	symbols      []string
	matches      func(a, b string) bool
	commentTypes []Pair
	stringTypes  []Pair
}

func NewAnalyzer(symbols, keywords []string, stringTypes, commentTypes []Pair, caseSensitive bool) LexAnalyzer {
	var a LexAnalyzer
	a.symbols = symbols
	a.keywords = keywords
	a.stringTypes = stringTypes
	a.commentTypes = commentTypes
	if caseSensitive {
		a.matches = func(a, b string) bool {
			return a == b
		}
	} else {
		a.matches = func(a, b string) bool {
			a = strings.ToUpper(a)
			b = strings.ToUpper(b)
			return a == b
		}
	}
	return a
}

func (me *LexAnalyzer) isComment(val string, token, closer *string) bool {
	for _, v := range me.commentTypes {
		if v.opens(val) {
			*token = v.parse(val)
			*closer = v.Closer
			return true
		}
	}
	return false
}

func (me *LexAnalyzer) isString(val string, token, closer *string) bool {
	for _, v := range me.stringTypes {
		if v.opens(val) {
			*token = v.parse(val)
			*closer = v.Closer
			return true
		}
	}
	return false
}

//Indicates whether or not the given value is a keyword, and if it is, it adjusts for casing.
func (me *LexAnalyzer) isKeyword(val string) (string, bool) {
	for _, kw := range me.keywords {
		if me.matches(val, kw) {
			return kw, true
		}
	}
	return "", false
}

//Indicates whether or not the given value is a symbol.
func (me *LexAnalyzer) isSymbol(val string, point int) bool {
	for _, kw := range me.symbols {
		if pt2 := point + len(kw); pt2 <= len(val) {
			if val[point:pt2] == kw {
				return true
			}
		}
	}
	return false
}

func (me *LexAnalyzer) getSymbol(val string, point int) string {
	for _, kw := range me.symbols {
		if pt2 := point + len(kw); pt2 <= len(val) {
			if val[point:pt2] == kw {
				return kw
			}
		}
	}
	return ""
}

//Returns: the token type, the token, the point in the file where the tokenizer left off
func (me *LexAnalyzer) NextToken(val string, point int) (int, string, int) {
	var token, closer string
	switch {
	case isWhitespace(val[point]):
		return Whitespace, string(val[point]), point + 1
	case isCharacter(val[point]): //is a keyword or identifier
		for !me.isSymbol(val, point) && !isWhitespace(val[point]) {
			token += string(val[point])
			point++
		}
		if kw, b := me.isKeyword(token); b { //is keyword
			return Keyword, kw, point
		}
		//is identifier
		found := false
		for _, v := range me.identTable {
			if v == token {
				found = true
				break
			}
		}
		if !found {
			me.identTable = append(me.identTable, token)
		}
		return Identifier, token, point
	case me.isComment(val[point:], &token, &closer):
		point += len(token) + len(closer)
		return Comment, token, point
	case me.isSymbol(val, point):
		s := me.getSymbol(val, point)
		return Symbol, s, point + len(s)
	default: //is a literal
		if isNumber(val[point]) { //is a number literal
			for ; point < len(val) && isNumber(val[point]); point++ {
				token += string(val[point])
			}
			found := false
			for _, v := range me.numLitTable {
				if v == token {
					found = true
					break
				}
			}
			if !found {
				me.numLitTable = append(me.numLitTable, token)
			}
			return IntLiteral, token, point
		} else if me.isString(val[point:], &token, &closer) {
			point += len(token) + len(closer)
			return StringLiteral, token, point
		}
	}

	return Error, string(val[point]), point
}
