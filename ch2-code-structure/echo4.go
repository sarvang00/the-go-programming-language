package main

import (
	"flag"
	"fmt"
	"strings"
)

// A flag to omit newline
var n = flag.Bool("n", false, "omit trailing newline")

// A seperator flag; default to " "
var sep = flag.String("s", " ", "separator")

func main() {
	// first, parse for flags
	flag.Parse()
	// Format the string as per the value of seperator flag
	fmt.Print(strings.Join(flag.Args(), *sep))
	// Add a newline if flag n is not added
	if !*n {
		fmt.Println()
	}
}
