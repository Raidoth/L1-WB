package main

import (
	"fmt"
)

/*
16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

/* CASE 1 */

//sort with go function

// func TQuickSort(arr []int) {
// 	sort.Ints(arr)
// }

/* CASE 2 */

func TQuickSort(arr []int) {
	//if < 2 len return
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	// Anything less than right will be put in left
	// and go to the next index, without reaching the right
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	TQuickSort(arr[:left])
	TQuickSort(arr[left+1:])
}

/* CASE 3 */
// func TQuickSort(arr []int) {
// 	const M = 12
// 	var i, j, l, r int
// 	var x int
// 	var low, high = make([]int, 0, M), make([]int, 0, M)

// 	low, high = append(low, 0), append(high, len(arr)-1)
// 	for { // (*take top request from stack*)
// 		l, low = low[len(low)-1], low[:len(low)-1]
// 		r, high = high[len(high)-1], high[:len(high)-1]
// 		for { // (*partition a[l] ... a[r]*)
// 			i, j = l, r
// 			x = arr[l+(r-l)/2]
// 			for {
// 				for ; arr[i] < x; i++ {
// 				}
// 				for ; x < arr[j]; j-- {
// 				}
// 				if i <= j {
// 					arr[i], arr[j] = arr[j], arr[i]
// 					i++
// 					j--
// 				}
// 				if i > j {
// 					break
// 				}
// 			}
// 			if i < r { // (*stack request to sort right partition*)
// 				low, high = append(low, i), append(high, r)
// 			}
// 			r = j // (*now l and r delimit the left partition*)
// 			if l >= r {
// 				break
// 			}
// 		}
// 		if len(low)+len(high) == 0 {
// 			break
// 		}
// 	}

// }

func main() {

	arr := []int{5, 234, 41, 32523, -4, 2, 55, 12, 64, 23, 22, -32}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	TQuickSort(arr)
	fmt.Println("After sort")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
