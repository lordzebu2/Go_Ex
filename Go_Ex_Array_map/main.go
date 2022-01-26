package main

import "fmt"


func map_ex() {
	js := map[string]string{"name":"nico","age":"20"}
	//map의 key값은 string vlaue값도 string
	for key, value := range js {
		fmt.Println(key, value, "map에 range")
	}
	fmt.Println(js)
}


func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "jjs")
	//배열에 값 넣는방법
	fmt.Println(names)
	map_ex()
}