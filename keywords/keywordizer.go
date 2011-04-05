package main

import (
	"strings"
	"fmt"
	"flag"
)

func block(val string) string {
	return "[" + val + strings.ToUpper(val) + "]"
}

type Keyword string

func (me Keyword) out() string {
	out := ""
	for i := 0; i < len(me); i++ {
		out += block(string(me[i]))
	}
	out += "\t {return(" + strings.ToUpper(string(me)) + ");}"
	return out
}

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(Keyword(strings.ToLower(flag.Arg(i))).out())
	}
}
