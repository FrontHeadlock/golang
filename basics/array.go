package main

import "fmt"

func main() {
	var t [5]int16 = [5]int16{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Print(t[i])
	}
	fmt.Println()

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
