package main

import (
	"fmt"
	"sync"
)

/*
3.	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений
*/

/* CASE 1 */

func main() {
	//create synchronization primitives
	var wg sync.WaitGroup
	var mtx sync.Mutex
	//create array with variable
	numbers := [...]int{2, 4, 6, 8, 10}
	//create variable with sum
	var sum int
	//create amount goroutine which equal slice length
	for _, val := range numbers {
		//add task in synchronization primitive
		wg.Add(1)
		go func(num int) {
			//lock the sum for other threads
			mtx.Lock()
			//add square number in sum
			sum += num * num
			//unlock the sum for other threads
			mtx.Unlock()
			defer wg.Done()
		}(val)

	}
	//waiting completion all tasks
	wg.Wait()
	//output sum
	fmt.Println(sum) //220
}

/* CASE 2 */

// func main() {
// 	//create array with variable
// 	numbers := [...]int{2, 4, 6, 8, 10}
// 	//create variable with sum
// 	var sum int
// 	//create channel int
// 	ch := make(chan int)
// 	//create amount goroutine which equal slice length
// 	for _, val := range numbers {
// 		go func(num int) {
// 			//put square number in channel
// 			ch <- int(math.Pow(float64(num), 2))
// 		}(val)
// 		//get number from channel and add in sum
// 		sum += <-ch
// 	}
// 	//output sum
// 	fmt.Println(sum) //220
// }
