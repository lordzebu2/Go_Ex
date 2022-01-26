package main

import "fmt"

func forLoop(){

	x:=0

	for i:=1; i<=100; i++{
		x = x+i
	}

	var(
		a, b, c = 1, 2, "안녕"
	)
 
	_=b
	_=c

	xx:=2e-3

	Loop:
		for i:=0; i<3; i++{
			for j:=0; j<3; j++{
				if j == 2 {
					continue Loop
				}
				fmt.Println(i, j)
			}
		}
		fmt.Println("Loop종료")
		
	for i, j := 0,0; i<5; i, j=i+1, j+2{
		fmt.Println(i, j)
	}
	fmt.Println("2변수 반복 종료")


	fmt.Println(a)
	fmt.Println("Hello")
	fmt.Println(x)
	fmt.Println(xx)
	
}