package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for index, number := range numbers{
		//range는 index를 준다. range를 이용한 for문.
		fmt.Println(index, number)
	}
	return 1
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		//go의 if는 조건 전에 변수생성 가능. if안에서만 쓸수있는것
		return false
	}
	return true
}

func main()  {
	superAdd(1,2,3,4,5,6)
	fmt.Println(canIDrink(15))
	fmt.Println(canIDrink_Switch(18), "스위치")
}