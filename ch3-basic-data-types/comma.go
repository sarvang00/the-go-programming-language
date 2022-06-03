package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	str := bufio.NewReader(os.Stdin)
	ip_str, _ := str.ReadString('\n')
	ip_str = strings.TrimSpace(ip_str)
	var rem string = ""
	idx := strings.Index(ip_str, ".")
	if idx != -1 {
		rem = ip_str[idx:]
		ip_str = ip_str[:idx]
	}
	fmt.Println(comma(ip_str) + rem)
}
