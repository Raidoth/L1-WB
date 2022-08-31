package main

import (
	"fmt"
	"strings"
)

/*
26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false


*/

/* CASE 1 */

func CheckUnicSymb(str string) bool {

	str = strings.ToLower(str)
	symbArr := make([]rune, 5)
	//iteration on string and symbArr, if val in string equal val in symbArr, then return false
	for _, s := range str {
		for _, val := range symbArr {
			if val == s {
				return false
			}
		}
		symbArr = append(symbArr, s)
	}

	return true
}

/* CASE 2 */

// func CheckUnicSymb(str string) bool {
// 	str = strings.ToLower(str)
// 	//iterations on string and string+1 for no check self
// 	for i := 0; i < len(str); i++ {
// 		for j := i + 1; j < len(str); j++ {
// 			if str[i] == str[j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true

// }

/* CASE 3 */

// type void struct{}

// type set map[rune]void

// //func add set
// func (s *set) addSet(symb rune) bool {
// 	if _, ok := (*s)[symb]; ok {
// 		return false
// 	}
// 	(*s)[symb] = void{}
// 	return true
// }

// func CheckUnicSymb(str string) bool {
// 	str = strings.ToLower(str)
// 	//create set
// 	symbSet := make(set, 5)
// 	//iteration on set, if symb find return false
// 	for _, val := range str {
// 		if ok := symbSet.addSet(val); !ok {
// 			return false
// 		}
// 	}
// 	return true
// }

/* CASE 4 */

// const UpperCaseToDown = 32

// func CheckUnicSymb(str string) bool {
// 	symbMap := make(map[byte]int, 5)
// 	//itaration on string
// 	for i := range str {
// 		s := str[i]
// 		//check symb have upper symb
// 		if s >= 65 && s <= 90 {
// 			s += UpperCaseToDown
// 		}
// 		//counter symb
// 		symbMap[s]++
// 		//if have more 1 symb return false
// 		if symbMap[s] > 1 {
// 			return false
// 		}
// 	}

// 	return true
// }

func main() {
	fmt.Println(CheckUnicSymb("abcd"))
	fmt.Println(CheckUnicSymb("abCdefAaf"))
	fmt.Println(CheckUnicSymb("aabcd"))
}
