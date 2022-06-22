package tempconv

import (
	"fmt"
	"os"
	"strconv"
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Fahrenheit temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func UnitConversion(val string) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := Fahrenheit(t)
	c := Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, FToC(f), c, CToF(c))
}
