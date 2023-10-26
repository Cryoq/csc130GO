package main

import (
	"fmt"
)

func calcNet(inc float32, taxes float32) float32 {
	var tax float32 = (100 - taxes) / 100
	return (tax * float32(inc))
}

func main() {
	fmt.Printf("Please enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Hello %v, What is your gross annual income? ", name)
	var income float32
	fmt.Scan(&income)

	fmt.Print("What is the percentage tax rate in your location? ")
	var tax float32
	fmt.Scan(&tax)

	var netIncome float32 = calcNet(income, tax)
	fmt.Printf("Well %v, that means that your net income is $%v", name, netIncome)
}
