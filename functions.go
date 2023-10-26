package main

import (
	"fmt"
	"math"
)

func info(x string) string {
	var name string
	fmt.Printf("What is the name of Country #%v: ", x)
	fmt.Scan(&name)
	return name
}

func currentPopulation(x string) float64 {
	var number float64
	fmt.Printf("What is the current population of %v? ", x)
	fmt.Scan(&number)
	if number < 0 {
		fmt.Println("That doesn't seem right. Please enter a positive number")
		currentPopulation(x)
	}
	return number
}

func annualGrowth(x string) float64 {
	var number float64
	fmt.Printf("What is the annual population growth rate of %v? ", x)
	fmt.Scan(&number)
	if number < -5 || number > 10 {
		fmt.Println("That doesn't seem right. Please enter a positive number")
		annualGrowth(x)
	}
	return number
}

func years() (uint, uint) {
	var years, interval uint
	fmt.Print("How many years of comparison should the table show? ")
	fmt.Scan(&years)
	for years < 0 {
		fmt.Println("That doesn't seem right. Please enter a value >= 1")
		fmt.Scan(&years)
	}
	fmt.Print("How many years should the intervals be? ")
	fmt.Scan(&interval)
	for interval < 0 {
		fmt.Println("That doesn't seem right. Please enter a value >= 1")
		fmt.Scan(&interval)
	}
	return years, interval
}

func calculatePopulation(p, r, t float64) uint {
	var pop float64 = p * (math.Pow((1 + (r / 100)), t))
	var popint uint = uint(pop)
	return popint
}

func header(country1, country2 string) {
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Printf("Years\t%v\t\t%v", country1, country2)
	fmt.Println()
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func table(pop1, pop2, rate1, rate2, years, interval float64) {
	var currentYear float64
	for years >= currentYear {
		fmt.Printf("%v\t%v\t\t%v\n", currentYear, calculatePopulation(pop1, rate1, currentYear), calculatePopulation(pop2, rate2, currentYear))
		currentYear += interval
	}
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
}

func main() {
	var count1, count2 string
	var pop1, pop2, rate1, rate2 float64
	var year, inter uint

	count1 = info("1")
	count2 = info("2")
	pop1 = currentPopulation(count1)
	pop2 = currentPopulation(count2)
	rate1 = annualGrowth(count1)
	rate2 = annualGrowth(count1)
	year, inter = years()

	var yearF float64 = float64(year)
	var interF float64 = float64(inter)

	header(count1, count2)
	table(pop1, pop2, rate1, rate2, yearF, interF)
}
