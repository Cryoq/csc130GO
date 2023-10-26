package main

import (
	"fmt"
	"math"
)

func name() string {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	return name
}

func str(info string) float64 {
	var x float64
	fmt.Printf("Please enter the %v: ", info)
	fmt.Scan(&x)
	return x
}

func total(principal float64, rate float64, year float64) float64 {
	var compound float64 = principal * math.Pow((1+(rate/100)), year)
	return compound
}

func main() {
	var name string = name()
	var principal float64 = str("principal")
	var rate float64 = str("annual percentage rate")
	var years float64 = str("number of years")

	var compoundRate float64 = total(principal, rate, years)

	fmt.Printf("Hello %v, the final amount after %v years at %v%% is $%f", name, years, rate, compoundRate)

}
