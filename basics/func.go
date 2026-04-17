package main

import "fmt"

func main() {
	x, y := f()
	println(x, y)

	fmt.Println(add(1, 2, 3))

	// 클로저
	kkk := func(x, y int) int {
		return x + y
	}

	fmt.Println(kkk(1, 1))

	//지연
	defer two()
	one()

	// recover : runtime 패닉 중단하고, panic 호출에 전달된 값 반환
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	// 패닉 : 런타임오류 raise
	panic("PANIC")
}

// 다중값 반환
func f() (int, int) {
	return 5, 6
}

// 가변함수
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// 재귀함수
func fact(x uint) uint {
	if x == 0 {
		return 0
	}
	return x * fact(x-1)
}

func one() {
	fmt.Println("one")
}
func two() {
	fmt.Println("two")
}
