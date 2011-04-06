package main

import (
	"strings"
)

var identTable []string = make([]string, 0)
var numLitTable []string = make([]string, 0)
var keywords []string = []string{"int", "void", "print", "println", "function", "program", "true", "false",
	"if", "else", "begin", "end", "while", "return", "boolean", "null"}
var symbols []string = []string{".", ";", "=", ",", "(", ")", "==", "||", "&&", "!", "/", "*", "-", "+", "<<", ">>", "=<", "=>", "<", ">", "!="}

const (
	keyword    string = "keyword"
	identifier = "identifier"
	symbol     = "symbol"
	literal    = "literal"
	comment    = "comment"
	whitespace = "whitespace"
	error      = "error"
)

func match(a, b string) bool {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	return a == b
}

func isCharacter(val byte) bool {
	return (val < 91 && val > 64) || (val < 123 && val > 96)
}

func isNumber(val byte) bool {
	return (47 < val && val < 58)
}

//Indicates whether or not the given value is a keyword, and if it is, it adjusts for casing.
func isKeyword(val string) (string, bool) {
	for _, kw := range keywords {
		if match(val, kw) {
			return kw, true
		}
	}
	return "", false
}

//Indicates whether or not the given value is a symbol.
func isSymbol(val string, point int) bool {
	for _, kw := range symbols {
		if pt2 := point + len(kw); pt2 <= len(val) {
			if val[point:pt2] == kw {
				return true
			}
		}
	}
	return false
}

func isComment(val string, point int) bool {
	return point+1 < len(val) && val[point] == '/' && val[point+1] == '/'
}

func isWhitespace(val byte) bool {
	return val == ' ' || val == '\t' || val == '\n'
}

func getSymbol(val string, point int) string {
	for _, kw := range symbols {
		if pt2 := point + len(kw); pt2 <= len(val) {
			if val[point:pt2] == kw {
				return kw
			}
		}
	}
	return ""
}

//Returns: the token type, the token, the point in the file where the tokenizer left off
func nextToken(val string, point int) (string, string, int) {
	switch {
	case isWhitespace(val[point]):
		return whitespace, string(val[point]), point + 1
	case isCharacter(val[point]): //is a keyword or identifier
		token := ""
		for !isSymbol(val, point) && !isWhitespace(val[point]) {
			token += string(val[point])
			point++
		}
		if kw, b := isKeyword(token); b { //is keyword
			return keyword, kw, point
		}
		//is identifier
		found := false
		for _, v := range identTable {
			if v == token {
				found = true
				break
			}
		}
		if !found {
			identTable = append(identTable, token)
		}
		return identifier, token, point
	case isComment(val, point):
		token := ""
		for ; val[point] != '\n'; point++ {
			token += string(val[point])
		}
		return comment, token, point
	case isSymbol(val, point):
		s := getSymbol(val, point)
		return symbol, s, point + len(s)
	default: //is a literal
		token := ""
		if isNumber(val[point]) { //is a number literal
			for ; isNumber(val[point]); point++ {
				token += string(val[point])
			}
			found := false
			for _, v := range numLitTable {
				if v == token {
					found = true
					break
				}
			}
			if !found {
				numLitTable = append(numLitTable, token)
			}
		} else {
			point++
			for ; val[point] != '"'; point++ {
				token += string(val[point])
			}
			point++
		}
		return literal, token, point
	}

	return error, string(val[point]), point
}
