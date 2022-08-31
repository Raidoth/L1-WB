package main

import "fmt"

/*
12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
/* CASE 1 */

type void struct{}
type set_t map[string]void

/* CASE 2 */

//type set_t map[string]bool

func main() {
	//create array
	str := [...]string{"cat", "cat", "dog", "cat", "tree"}
	//create set
	set := make(set_t, 5)
	//put string in set
	for _, v := range str {
		//set[v] = true
		set[v] = void{}
	}
	//output set
	for val := range set {
		fmt.Println(val)
	}
}
