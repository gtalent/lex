package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for i, a := range args {
		fmt.Println(i, a)
	}
}

func test1() {
	fmt.Println("Test 1:")
	fmt.Println("\t", !isSymbol("rri", 0))
	fmt.Println("\t", isSymbol(">=", 0))
	fmt.Println("\t", isSymbol("!=", 0))
}

func test2() {
	fmt.Println("Test 2:")
	fmt.Println("\t", isWhitespace(' '))
	fmt.Println("\t", isWhitespace('\n'))
	fmt.Println("\t", isWhitespace('\t'))
}
