package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	invalid := 0 // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	r, err := in.ReadString('\n') // returns rune, error
	r = strings.TrimSpace(r)
	byteData := []rune(r)
	for rDat := 0; rDat < len(byteData); rDat++ {
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if byteData[rDat] == unicode.ReplacementChar {
			invalid++
			continue
		}
		counts[byteData[rDat]]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
