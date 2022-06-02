package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Feet float64
type Meter float64

func main() {
	// Starting with feet
	fmt.Printf("Enter the length in feet: ")
	feetReader := bufio.NewReader(os.Stdin)
	feetStr, _ := feetReader.ReadString('\n')
	feetStr = strings.TrimSpace(feetStr)
	f, _ := strconv.ParseFloat(feetStr, 64)
	feet := Feet(f)
	fmt.Printf("Length: %s or %s\n", feet, feetToMeter(feet))

	// Ending with meter
	fmt.Printf("Enter the length in meter: ")
	meterReader := bufio.NewReader(os.Stdin)
	meterStr, _ := meterReader.ReadString('\n')
	meterStr = strings.TrimSpace(meterStr)
	m, _ := strconv.ParseFloat(meterStr, 64)
	meter := Meter(m)
	fmt.Printf("Length: %s or %s\n", meter, meterToFeet(meter))
}

func feetToMeter(f Feet) Meter {
	return Meter(f * 10000 / 3048)
}

func meterToFeet(m Meter) Feet {
	return Feet(m * 3048 / 10000)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft.", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm.", m)
}
