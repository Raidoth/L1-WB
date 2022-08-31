package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*

4.	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

//create synchronization primitive
var wg sync.WaitGroup

func Worker(id int, inputCh <-chan string) {
	//output info about ends worker
	defer fmt.Println("Worker end ", id)
	//infinity cycle gets string from channel
	for {
		//if string is received, outputs id goroutine and string, otherwise remove task in synchronization primitive and return
		if res, ok := <-inputCh; ok {
			fmt.Printf("Goroutine %d, send %s\n", id, res)
		} else {
			wg.Done()
			return
		}
	}

}

func main() {
	//create variable amount workers
	var n int
	fmt.Println("Input amount workers")
	//add amount in variable which input user
	fmt.Scan(&n)
	if n == 0 {
		return
	}
	//create channel for data
	var inputCh = make(chan string)
	//create amount goroutine which equal n(input user)
	for i := 0; i < n; i++ {
		//add task in synchronization primitive
		wg.Add(1)
		go Worker(i, inputCh)
	}
	//create channel for exit
	quit := make(chan os.Signal, 1)
	//channel wait ctrl+c
	signal.Notify(quit, os.Interrupt)
	//infinity cycle send string in channel, if user send ctrl+c channel close and wait complete all task
	for {
		select {
		case <-quit:
			close(inputCh)
			close(quit)
			wg.Wait()
			fmt.Println("Program end")
			os.Exit(0)
		default:
			inputCh <- RandomString(5)
			time.Sleep(time.Second)

		}

	}

}

//create random string with given length
func RandomString(cntSymb int) string {
	//create slice with all letters
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//create result slice given length
	s := make([]rune, cntSymb)
	//receiving random letters from slice and add result slice
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	//return string
	return string(s)
}
