package main

import "fmt"

/*
23.	Удалить i-ый элемент из слайса.
*/

/* CASE 1 */

func deleteSlice(str []string, elem int) []string {
	// create slice for new string
	newStr := make([]string, 0, len(str)-1)
	//iteration on string and if removable element find, skip him
	for i, val := range str {
		if i == elem {
			continue
		}
		newStr = append(newStr, val)
	}
	//return new slice
	return newStr
}

/* CASE 2 */

// func deleteSlice(str []string, elem int) []string {
// 	//add slice str["ASDASD", "SDAASDASDASDASD"] to str["ABC", "BCD", "CDE"]
// 	return append(str[:elem], str[elem+1:]...)
// }

/* CASE 3*/

/*
	copy string from elem+1 in elem
	string "ABC", "BCD", "CDE", "CCC", "ASDASD", "SDAASDASDASDASD"
	elem 3
	src str["ASDASD", "SDAASDASDASDASD"]
	dst str["CCC", "ASDASD", "SDAASDASDASDASD"]
	and return str len -1
*/

// func deleteSlice(str []string, elem int) []string {

// 	copy(str[elem:], str[elem+1:])
// 	return str[:len(str)-1]
// }

func main() {
	//create string slice
	str := []string{"ABC", "BCD", "CDE", "CCC", "ASDASD", "SDAASDASDASDASD"}
	fmt.Println("Before delete")
	//output slice before delete
	for _, val := range str {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
	fmt.Println("After delete")
	str = deleteSlice(str, 3)
	//output slice after delete
	for _, val := range str {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}
