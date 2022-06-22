package main

import (
	"bufio"
	"fmt"
	"main/Ch02/exer01"
	"os"
)

// Write a general-purpose unit-conversion program analogous to cf that reads numbers from
// its command-line arguments or from the standard input if there are no arguments, and converts
// each number into units like temperature in Celsius and Fahrenheit, length in feet and meters,
// weight in pounds and kilograms, and the like.
func main() {
	if len(os.Args) > 1 {
		fromCmdArgs()
	} else {
		fromStdin()
	}
}
func fromCmdArgs() {
	args := os.Args[1:]
	for _, arg := range args {
		tempconv.UnitConversion(arg)
	}
}

func fromStdin() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		tempconv.UnitConversion(s.Text())
	}
	if s.Err() != nil {
		fmt.Printf("Error: %v\n", s.Err())
		os.Exit(1)
	}
}
