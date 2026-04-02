package main

import "fmt"

func main() {
	// Integer + integer
	fmt.Println("1+1 =", 1+1)
	// Float + float
	fmt.Println("1+1 =", 1.0+1.0)
	// Mixed numeric constants (둘 다 상수 리터럴이라 가능)
	fmt.Println("1+1 =", 1.0+1)

	// Boolean AND examples
	fmt.Println(true && true)
	fmt.Println(true && false)
	// Boolean OR examples
	fmt.Println(true || true)
	fmt.Println(true || false)
	// Boolean NOT example
	fmt.Println(!true)
}
