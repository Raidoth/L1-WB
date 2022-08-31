package main

import (
	"fmt"
)

/*
8.	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

/*
15 - ...1111
change 3 bit
if bit on pos equal 0, then 1 xor 0 equal 1
if bit on pos equal 1, then 1 xor 1 equal 0
*/

func ChangeBit(num *int64, pos uint) {
	*num ^= (1 << pos)
}

func main() {

	var num int64
	num = 15
	fmt.Printf("%.5d bytes %.64b\n", num, num)
	ChangeBit(&num, 3)
	fmt.Printf("%.5d bytes %.64b\n", num, num)
}
