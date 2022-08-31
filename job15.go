package main

import (
	"fmt"
	"math/rand"
)

/*
15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string //global variable and unusable var
func someFunc() {
	v := createHugeString(1 << 10) // no function, also string can return string < 2^10
	justString = v[:100]	//panic if v<100

}
func main() {
	someFunc()
}
*/

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}

func main() {
	someFunc()
}

//create random string with given length
func createHugeString(cntSymb int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, cntSymb)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
