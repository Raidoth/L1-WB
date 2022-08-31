package main

import (
	"fmt"
)

/*
9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	//create array with num
	numbers := [...]int{2, 3, 12, 43, 14, 5, 6, 22}
	//create channel for num
	channel1 := make(chan int)
	//create channel for num*2
	channel2 := make(chan int)
	//goroutine put number from array in channel1
	go func() {

		for _, val := range numbers {
			channel1 <- val
		}
		close(channel1)

	}()
	//goroutine gets number from channel1 , multiply 2 and put in channel2
	go func() {
		for val := range channel1 {
			channel2 <- val * 2
		}
		close(channel2)

	}()
	//gets number from channel2 and put in stdout
	for val := range channel2 {
		fmt.Println(val)
	}

}
