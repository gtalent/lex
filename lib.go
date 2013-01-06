/*
   Copyright 2011-2012 gtalent2@gmail.com

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

func Tokens(input string) []Token {
	var tokens []Token

	symbols := []string{"&&", "||", "=<", "=>", "==", "!=", "<", ">", "/", "*", "-", "+", "(", ")", "!"}
	commentTypes := []Pair{{Opener: "#", Closer: "\n"}}
	lex := NewAnalyzer(symbols, []string{}, commentTypes, true)
	for point := 0; point < len(input); {
		var t Token
		t.TokType, t.TokValue, point = lex.nextToken(input, point)
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

func (me *Pair) isComment(val string) bool {
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
}

func NewAnalyzer(symbols, keywords []string, commentTypes []Pair, caseSensitive bool) LexAnalyzer {
	var a LexAnalyzer
	a.symbols = symbols
	a.keywords = keywords
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

func (me *LexAnalyzer) isComment(val string, point int, opener, closer *string) bool {
	val = val[point:]
	for _, v := range me.commentTypes {
		if v.isComment(val) {
			*opener = v.Opener
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
func (me *LexAnalyzer) nextToken(val string, point int) (int, string, int) {
	var opener, closer string
	switch {
	case isWhitespace(val[point]):
		return Whitespace, string(val[point]), point + 1
	case isCharacter(val[point]): //is a keyword or identifier
		token := ""
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
	case me.isComment(val, point, &opener, &closer):
		token := ""
		point += len(opener)
		for ; val[:len(closer)] != closer; point++ {
			val = val[1:]
			token += string(val[0])
		}
		return Comment, token, point
	case me.isSymbol(val, point):
		s := me.getSymbol(val, point)
		return Symbol, s, point + len(s)
	default: //is a literal
		token := ""
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
		} else {
			point++
			for ; val[point] != '"'; point++ {
				token += string(val[point])
			}
			point++
			return StringLiteral, token, point
		}
	}

	return Error, string(val[point]), point
}
