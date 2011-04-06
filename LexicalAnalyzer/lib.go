package main

import "strings"

var symbolTable []string = make([]string, 0)
var keywords []string = []string{"int", "void", "print", "println", "function", "program", "true", "false",
	"if", "else", "begin", "end", "while", "return", "boolean", "null"}
var symbols []string = []string{".", ";", "=", ",", "(", ")", "==", "||", "&&", "!", "/", "*", "-", "+", "<<", ">>", "=<", "=>", "<", ">", "!="}

func match(a, b string) bool {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	return a == b
}

func isCharacter(val byte) bool {
	return (val < 91 && val > 64) || (val < 123 && val > 96)
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
		if pt2 := point+len(kw); pt2 <= len(val) {
			if val[point:pt2] == kw {
				return true
			}
		}
	}
	return false
}

func isWhitespace(val byte) bool {
	return val == ' ' || val == '\t' || val == '\n'
}

func parseToken(val string, start int) (string, int) {
	token := ""
	point := start

	if isCharacter(val[start]) {
		//is a keyword or identifier
		for ; !isSymbol(val, point) && !isWhitespace(val[point]); point++ {
			token += string(val[point])
		}
	}

	return token, point
}
