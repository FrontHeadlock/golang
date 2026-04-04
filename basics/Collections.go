package main

import "fmt"

func main() {
	printArray()
	fmt.Println()
}

func printArray() {
	fmt.Println("[Arrays]")

	var numbers [5]int
	numbers[4] = 100
	fmt.Println("numbers =", numbers)
	fmt.Println("numbers[4] =", numbers[4])

	scores := [5]float64{98, 93, 77, 82, 83}
	var total float64

	for _, score := range scores {
		total += score
	}

	fmt.Println("scores length =", len(scores))
	fmt.Println("average =", total/float64(len(scores)))
}
