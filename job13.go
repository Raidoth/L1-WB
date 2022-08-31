package main

import "fmt"

/*
13.	Поменять местами два числа без создания временной переменной.
*/

func main() {

	a := 5
	b := 4
	fmt.Println("a = ", a, " b = ", b)
	/* CASE 1 */
	//change with help of language

	a, b = b, a
	fmt.Println("a = ", a, " b = ", b)

	/* CASE 2 */
	//change with help of maths

	// a += b    //a=9 b = 4
	// b = a - b //a=9 b=5
	// a -= b    //a=4 b=5
	// fmt.Println("a = ", a, " b = ", b)

	/* CASE 3 */
	//change with help of bool

	//a=101(5) b=100(4)

	// a ^= b    //a=101^100=001(1)
	// b = a ^ b //b=001^100=101(5)
	// a ^= b    //a=001^101=100(4)
	// fmt.Println("a = ", a, " b = ", b)

}
