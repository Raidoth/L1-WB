package main

import (
	"fmt"
	"sync"
)

/*
2.	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
/* CASE 1 */

func main() {
	//create array with variable
	numbers := [...]int{2, 4, 6, 8, 10}
	//create synchronization primitive
	var wg sync.WaitGroup
	//create amount goroutine which equal slice length
	for _, v := range numbers {
		//add task in synchronization primitive
		wg.Add(1)
		go func(num int) {
			//output square number which compute in printf
			fmt.Printf("Square %d equal %d\n", num, num*num)
			//Delete task in synchronization primitive
			wg.Done()
		}(v)

	}
	//waiting completion all tasks
	wg.Wait()
}

/* CASE 2 */

// func main() {
// 	//create array with variable
// 	numbers := [...]int{2, 4, 6, 8, 10}
// 	//create channel int
// 	ch := make(chan int)
// 	//create amount goroutine which equal slice length
// 	for _, v := range numbers {
// 		go func(num int) {
// 			//output square number which compute in printf
// 			fmt.Printf("Square %d equal %d\n", num, num*num)
// 			//put number in channel for wait execute all goroutines
// 			ch <- num
// 		}(v)

// 	}
// 	//wait all out numbers
// 	for i := 0; i < len(numbers); i++ {
// 		<-ch
// 	}

// }

/* CASE 3 */

// func main() {
// 	//create slice with variable
// 	numbers := [...]int{2, 4, 6, 8, 10}
// 	//create channel int
// 	ch := make(chan int)
// 	//create amount goroutine which equal slice length
// 	for _, v := range numbers {
// 		go func(num int) {
// 			//put square numbers in channel
// 			ch <- num * num
// 		}(v)
// 		//output numbers square in stdout
// 		fmt.Printf("Square %d equal %d\n", v, <-ch)
// 	}
// 	close(ch)

// }
