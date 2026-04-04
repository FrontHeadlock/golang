package main

import "fmt"

func main() {
	// 입력값을 저장할 공간을 먼저 선언
	fmt.Print("숫자를 입력하시오: ")
	var num float64
	fmt.Scanf("%f", &num)

	// 계산 결과를 새 변수에 바로 저장할 때는 := 사용
	doubled := num * 2
	fmt.Println("입력값 x 2 =", doubled)

	// 화씨를 입력받아 섭씨로 변환: C = (F - 32) * 5 / 9
	fmt.Print("화씨를 입력하시오: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("섭씨: %gC\n", celsius)

	// 피트를 미터로 변환 (1 ft = 0.3048 m)
	fmt.Print("피트를 입력하시오: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Printf("미터: %gm\n", meters)
}
