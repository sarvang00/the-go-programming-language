// Exercise 3.10
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ipStr, _ := reader.ReadString('\n')
	ipStr = strings.TrimSpace(ipStr)
	fmt.Println(noRecurComma(ipStr))
}

func noRecurComma(s string) string {
	var buf bytes.Buffer
	if len(s) > 3 {
		for i := 0; i < len(s); i++ {
			if i%3 == 0 && i > 0 {
				buf.WriteString(",")
			}
			// buf.WriteString(strconv.Itoa(i))
			buf.WriteString(s[i : i+1])
		}
	} else {
		buf.WriteString(s)
	}
	return buf.String()
}
