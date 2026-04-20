package main

import "fmt"

func main() {
	// 출력 함수
	// Print, Println , Printf(포맷)

	//var a int = 10
	//var b int = 20
	//var f float64 = 32783838.2218

	//fmt.Print("a:", a, "b:", b)
	//fmt.Println("a:", a, "b:", b)
	//fmt.Printf("a: %d, f: %fd\n", a, f)

	//입력 함수
	// Scan, Scanf, Scanln

	var a1 int
	var b1 int

	n, err := fmt.Scan(&a1, &b1)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a1, b1)
	}
}
