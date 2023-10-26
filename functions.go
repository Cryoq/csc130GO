package main

import (
	"fmt"
	"math"
)

func intro() {
	fmt.Println("This program will calculate the integral of the function\n\t3x^3 - 2x^2\nbetween user defined limits: a and b")
}

func data(letter string) float32 {
	var num float32
	fmt.Printf("What is the value of \"%v\": ", letter)
	fmt.Scan(&num)
	return num
}

func deltaInfo() {
	fmt.Println("The accuracy of this calculation depends on the value of delta that you use.")
}

func f(x32 float32) float32 {
	var x float64 = float64(x32)
	var value64 float64 = (3 * math.Pow(x, 3)) - (2 * math.Pow(x, 2))
	var value32 = float32(value64)
	return value32
}

func rimann(lower float32, upper float32, delta float32) float32 {
	var total float32
	for lower < upper {
		total += f(lower) * delta
		lower += delta
	}
	return total
}

func main() {
	intro()
	var a float32 = data("a")
	var b float32 = data("b")
	deltaInfo()
	var delta float32 = data("delta")

	var tot float32 = rimann(a, b, delta)
	fmt.Printf("The integral over the provided limits is %f", tot)

}
