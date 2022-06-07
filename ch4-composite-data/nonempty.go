// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.

func nonempty(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func swapChars(str1 *string, str2 *string) {
	tp := *str1
	*str1 = *str2
	*str2 = tp
}

// Ex. 4.3 - Rewrite reverse to use an array pointer instead of a slice
func revStr(ipStr []string) {
	for i, j := 0, len(ipStr)-1; i < j; i, j = i+1, j-1 {
		swapChars(&ipStr[i], &ipStr[j])
	}
}

// Ex. 4.4 - A version of rotate that operates in a single pass
func rotateStrs(ipStr []string) []string {
	tpStr := ipStr[0]
	for i := 0; i < len(ipStr)-1; i++ {
		ipStr[i] = ipStr[i+1]
	}
	ipStr[len(ipStr)-1] = tpStr
	return ipStr
}

func main() {
	data := []string{"one", "two", "three", "four"}
	data = nonempty(data)
	fmt.Println(data)
	revStr(data)
	fmt.Println(data)
	rotData := rotateStrs(data)
	fmt.Println(rotData)
}
