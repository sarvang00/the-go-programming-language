package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC)) // "180" °F

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	fmt.Println(c == Celsius(f)) // "true"!

	cel := FtoC(212.0)
	fmt.Println(cel.String()) // "100°C"
	fmt.Printf("%v\n", cel)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", cel)   // "100°C"
	fmt.Println(cel)          // "100°C"
	fmt.Printf("%g\n", cel)   // "100"; does not call String
	fmt.Println(float64(cel)) // "100"; does not call String
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
