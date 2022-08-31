package main

import "fmt"

/*
11.	Реализовать пересечение двух неупорядоченных множеств.
*/

//create set from map
type void struct{}
type set_t map[int]void

func main() {
	//create first set
	a := set_t{
		8:  void{},
		20: void{},
		10: void{},
		3:  void{},
		6:  void{},
		11: void{},
		15: void{},
		0:  void{},
	}
	//output firs set
	for key := range a {
		fmt.Printf("%d ", key)
	}
	fmt.Println()
	//create second set
	b := set_t{
		15:  void{},
		300: void{},
		1:   void{},
		3:   void{},
		5:   void{},
		20:  void{},
		10:  void{},
	}
	//output second set
	for key := range b {
		fmt.Printf("%d ", key)
	}
	fmt.Println()
	//create intersection set
	intersection := make(set_t, 5)
	//check less set and put intersection in set
	if len(a) < len(b) {
		for key := range a {
			if _, ok := b[key]; ok {
				intersection[key] = void{}
			}
		}
	} else {
		for key := range b {
			if _, ok := a[key]; ok {
				intersection[key] = void{}
			}
		}
	}
	//output intersection
	for key := range intersection {
		fmt.Printf("%d ", key)
	}
	fmt.Println()
}
