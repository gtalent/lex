package main

import (
	"strings"
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
	args := flag.Args()
	//for arg := range args {
	for i := 0; i < len(args); i++ {
		println(Keyword(args[i]).out())
	}
}
