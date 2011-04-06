package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse();
	args := flag.Args()
	for i, a := range args {
		fmt.Println(i, a)
	}
}

func isKeyword(val string) {

}
