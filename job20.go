package main

import (
	"fmt"
	"strings"
)

/*
20.	Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

/* CASE 1 */

func ReversWordsStr(s *string) string {
	//create slice words
	Words := make([]string, 0, 5)
	//create variable lastword pos
	var lastWord int
	//cycle on string and put in Words slice
	for i, val := range *s {

		if string(val) == " " {
			Words = append(Words, (*s)[lastWord:i])
			lastWord = i + 1
		}
	}
	//add last word
	Words = append(Words, (*s)[lastWord:])

	var sb strings.Builder
	//iterations on words slice and add in builder from the end
	for i := range Words {
		sb.WriteString(Words[len(Words)-1-i] + " ")

	}
	//return string
	return sb.String()

}

/* CASE 2 */

// func ReversWordsStr(s *string) string {
// 	//splits a string into words
// 	arrWord := strings.Fields(*s)
// 	//create string builder
// 	var sb strings.Builder
// 	//iterations on words slice and add in builder from the end
// 	for i := range arrWord {
// 		sb.WriteString(arrWord[len(arrWord)-1-i])
// 		sb.WriteString(" ")
// 	}
// 	//return string
// 	return sb.String()
// }

/* CASE 3 */

// func ReversWordsStr(s *string) string {
//	//splits a string into words
// 	arrWord := strings.Fields(*s)
// 	//change places words in array
// 	for i, j := 0, len(arrWord)-1; i < j; i, j = i+1, j-1 {
// 		arrWord[i], arrWord[j] = arrWord[j], arrWord[i]
// 	}
// 	//create string with separator
// 	return strings.Join(arrWord, " ")
// }

func main() {
	//create string
	s := "snow dog sun"
	//output string
	fmt.Println(s)
	res := ReversWordsStr(&s)
	//output rever string
	fmt.Println(res)

}
