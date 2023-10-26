package main

import (
	"fmt"
	"slices"
)

func header() {
	lines()
	fmt.Println("Welcome to Cyber Groceries")
	lines()
}

func lines() {
	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func midHeader(itemList []string, priceList []float32) {
	lines()
	fmt.Printf("Items = %v\n", itemList)
	fmt.Printf("Prices = %v\n", priceList)
	lines()
}

func numofItems() uint {
	var num uint
	fmt.Print("How many items is the customer buying? ")
	fmt.Scan(&num)
	return num
}

func items(num uint) []string {
	itemList := []string{}
	var item string
	for i := uint(1); i <= num; i++ {
		fmt.Printf("What is item number %v? ", i)
		fmt.Scan(&item)
		itemList = append(itemList, item)
	}
	return itemList
}

func price(list []string) []float32 {
	priceList := []float32{}
	var item float32
	for _, s := range list {
		fmt.Printf("What is the price of %v? ", s)
		fmt.Scan(&item)
		priceList = append(priceList, item)
	}
	return priceList
}

func maxIndex(list []float32) int {
	max := slices.Max(list)

	for i, v := range list {
		if v == max {
			return i
		}
	}
	return 0

}

func minIndex(list []float32) int {
	min := slices.Min(list)

	for i, v := range list {
		if v == min {
			return i
		}
	}
	return 0

}

func sum(list []float32) float32 {
	var sum float32 = 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	header()
	var num uint = numofItems()
	listofItems := items(num)
	listofPrices := price(listofItems)

	var maxI int = maxIndex(listofPrices)
	var minI int = minIndex(listofPrices)

	midHeader(listofItems, listofPrices)

	fmt.Printf("The cheapest item is %v\n", listofItems[minI])
	fmt.Printf("The most expensive item is %v\n", listofItems[maxI])

	var total float32 = sum(listofPrices)
	fmt.Printf("The total cost is %v\n", total)
	lines()
	lines()
}
