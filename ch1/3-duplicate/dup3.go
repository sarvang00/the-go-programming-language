package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ignoring err
		data, _ := ioutil.ReadFile(filename)
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		// if n > 1 {
		// 	fmt.Printf("%d\t%s\n", n, line)
		// }
		fmt.Println(line, n)
	}
}
