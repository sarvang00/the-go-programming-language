// exercise 4.1
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := sha256.Sum256([]byte("x"))
	y := sha256.Sum256([]byte("y"))
	mismatch_count := 0
	for i := 0; i < 32; i++ {
		if x[i] != y[i] {
			mismatch_count++
		}
	}
	fmt.Println("The number of mismatches are: ", mismatch_count)
}
