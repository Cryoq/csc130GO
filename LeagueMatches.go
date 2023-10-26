package main

import (
	"fmt"
)

func numTeams() uint16 {
	fmt.Print("How mant teams are in this league? ")
	var num uint16
	fmt.Scan(&num)
	return num
}

func matches(team uint16) uint16 {
	var total uint16
	for i := team; i > 0; i-- {
		total += (i - 1)
	}
	return total
}

func main() {
	var teams uint16 = numTeams()
	var numofMatches uint16 = matches(teams)
	fmt.Printf("A league of %v teams would require at least %v matches", teams, numofMatches)

}
