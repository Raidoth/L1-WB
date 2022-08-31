package main

import "fmt"

/*
1.	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
type Human struct {
	name   string
	height uint
	weight uint
}
type Action struct {
	Human
}

//method by Human struct
func (h *Human) ShowInfo() {
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", h.name, h.height, h.weight)
}

func main() {
	// created variable struct Action which inherit Human struct and add name "test" height and weight equal default 0
	a := Action{Human{name: "test"}}
	// add height and weight
	a.height = 150
	a.weight = 200
	//shows info about struct Action
	a.ShowInfo()
}
