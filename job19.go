package main

import (
	"fmt"
)

/*
19.	Разработать программу,
	которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

/* CASE 1 */
/*
create string res and iteration on input string
in cycle add to res input string and old res string
*/

func reverseString(s *string) string {
	var res string
	for _, s := range *s {
		res = string(s) + res
	}

	return res
}

/* CASE 2 */
/*
function create slice runes rn and slice runes from string res, and puts in res slice from the end string slice
*/

// func reverseString(s *string) string {

// 	rn := []rune(*s)
// 	var res []rune
// 	for i := range rn {
// 		res = append(res, rn[len(rn)-1-i])
// 	}
// 	return string(res)
// }

/* CASE 3 */
/*
function create builder and slice runes from string and add in builder from the end slice
*/

// func reverseString(s *string) string {

// 	var sb strings.Builder
// 	runes := []rune(*s)
// 	for i := range runes {
// 		sb.WriteRune(runes[len(runes)-i-1])
// 	}
// 	return sb.String()
// }

/* CASE 4*/
/*
function create runes from string and iteration on slice and change places runes
*/

// func reverseString(s *string) string {
// 	//create slice rune input string
// 	runes := []rune(*s)

// 	for left, right := 0, len(runes)-1; left < right-left; left++ {
// 		runes[left], runes[right-left] = runes[right-left], runes[left]
// 	}
// 	return string(runes)
// }

func main() {
	//create string
	s := "главрыба"
	//output string
	fmt.Println(s)
	//reverse string
	res := reverseString(&s)
	//output reverse string
	fmt.Println(res)
}
