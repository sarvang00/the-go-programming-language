package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Fibonacci
	fmt.Printf("Enter the number to get nth fibonacci: ")
	reader1 := bufio.NewReader(os.Stdin)
	str_out, _ := reader1.ReadString('\n')
	str_out = strings.TrimSpace(str_out)
	read_int, _ := strconv.Atoi(str_out)
	fmt.Printf("%d Fibonacci is %d\n", read_int, fib(read_int))

	// GCD
	fmt.Printf("Enter the first number for gcd: ")
	reader2 := bufio.NewReader(os.Stdin)
	x_str, _ := reader2.ReadString('\n')
	x_str = strings.TrimSpace(x_str)
	x, _ := strconv.Atoi(x_str)
	fmt.Printf("Enter the second number for gcd: ")
	reader3 := bufio.NewReader(os.Stdin)
	y_str, _ := reader3.ReadString('\n')
	y_str = strings.TrimSpace(y_str)
	y, _ := strconv.Atoi(y_str)
	fmt.Printf("GCD of %d and %d is %d\n", x, y, gcd(x, y))
}

// Calculating nth fibo
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

// Calculating gcd of 2 integers
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
