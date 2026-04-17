package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println("total:", total/float64(len(xs)))
}

func avg(xs []float64) float64 {
	// runtime 오류 raise 용
	panic("Not Implemented")
}
