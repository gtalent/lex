package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, a := range args {
		file, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Println(err.String())
			continue
		}
		val := string(file)
		point := 0
		var tokenType string
		for point < len(val) {
			tokenType, _, point = nextToken(val, point)
			fmt.Print(tokenType, ", ")
		}
	}
	/*
		test1()
		test2()
		test3()
	*/
}

//Tests symbol identifier.
func test1() {
	fmt.Println("Test 1:")
	fmt.Println("\tGood =", !isSymbol("rri", 0))
	fmt.Println("\tGood =", isSymbol(">=", 0))
	fmt.Println("\tGood =", isSymbol("!=", 0))
}

//Tests keyword identifier.
func test2() {
	fmt.Println("Test 2:")
	_, ok := isKeyword("function")
	fmt.Println("\tGood =", ok)
	_, ok = isKeyword("FUNCTION")
	fmt.Println("\tGood =", ok)
	_, ok = isKeyword("narf")
	fmt.Println("\tGood =", !ok)
}

//Tests whitespace identifier.
func test3() {
	fmt.Println("Test 3:")
	fmt.Println("\tGood =", isWhitespace(' '))
	fmt.Println("\tGood =", isWhitespace('\n'))
	fmt.Println("\tGood =", isWhitespace('\t'))
}
