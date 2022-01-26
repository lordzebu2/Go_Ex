package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20 //b는 a를 살펴보는 pointer이다.. a까지 값을 바꿔버림 주소값이 연결되어있음
	fmt.Println(a, b)
	fmt.Println(*b)
	fmt.Println(a)
}

//&는 주소값 보는것
//*는 주소값을 원래대로 보는것?

//&는 주소와 같은것 
//*는 살펴보는것 or *을 주소에 쓰면 주소에 담긴값을 변경 가능