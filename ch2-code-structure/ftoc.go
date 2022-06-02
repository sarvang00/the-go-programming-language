package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, fToC(boilingF))
	fmt.Printf("freezing point = %g°F or %g°C\n", freezingF, fToC(freezingF))
}

func fToC(f float64) float64 {
	return ((f - 32) * 5 / 9)
}
