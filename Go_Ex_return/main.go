package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func lenAndUpper_2(name string) (lenght int, uppercase string){
	//리턴할 것을 미리 명시해 준다 그럼 리턴할것을 안정해줘도 명시한것을 반환함

	defer fmt.Println("return완료")//defer는 function이 끝난 후에 실행되게 하는것

	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string){
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	//리턴방법

	totalLength, upperName := lenAndUpper("jjs")
	fmt.Println(totalLength, upperName)
	//2개 이상 리턴방법

	totalLength_2, upperName_2 := lenAndUpper_2("kdy") //lenAndUpper_2 func끝 defer발동
	fmt.Println(totalLength_2, upperName_2)
	//naked return 방법

	repeatMe("nico", "재성", "냐옹이", "AY", "이건뭐")
	//여러개의 argument 전달. argument=매개변수에 넣는 값, 전달값, 전달인자
}