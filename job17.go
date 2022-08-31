package main

import (
	"fmt"
	"sort"
)

/*
17.	Реализовать бинарный поиск встроенными методами языка
*/
/*
 	Begin with the mid element of the whole array as a search key.
    If the value of the search key is equal to the item then return an index of the search key.
    Or if the value of the search key is less than the item in the middle of the interval, narrow the interval to the lower half.
    Otherwise, narrow it to the upper half.
    Repeatedly check from the second point until the value is found or the interval is empty.
*/

func BinarySearch(_arr []int, val int) (int, bool) {
	left, right := 0, len(_arr)-1
	var middle int

	for left <= right {

		middle = (left + right) >> 1
		switch {

		case _arr[middle] < val:
			left = middle + 1
		case _arr[middle] > val:
			right = middle - 1
		case _arr[middle] == val:
			return middle, true

		}
	}
	return 0, false

}

func main() {
	//create slice int
	arr := []int{5, 234, 41, 32523, 2, 55, 12, 64, 23, 22}
	//sort slice for binarysearch
	sort.Ints(arr)
	//output all array
	for i, val := range arr {
		fmt.Printf("index %d, value %d\n", i, val)
	}
	findNum := 5
	val, ok := BinarySearch(arr, findNum)
	if ok {
		//out index find num
		fmt.Println("Find num", findNum, "Index find ", val)
	} else {
		fmt.Println("No found")
	}
}
