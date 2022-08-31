package main

import "fmt"

/*
14.	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
func checkType(data interface{}) string {
	switch data.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	}
	return "unknow type"
}

func main() {

	s := make([]interface{}, 0, 5)
	s = append(s, 5)
	s = append(s, true)
	ch := make(chan int)
	s = append(s, ch)
	s = append(s, "asdasd")
	s = append(s, 6.05)

	for _, val := range s {
		fmt.Println(checkType(val))
	}

}
