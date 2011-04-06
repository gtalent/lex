package main

import "strings"

func match(a, b string) bool {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	return a == b
}
