package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
5.	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

func main() {
	//create var is duration programm
	var duration int
	fmt.Println("Input duration in seconds:")
	//get user input duration
	fmt.Scan(&duration)
	if duration == 0 {
		return
	}
	//channel waiting diration and out
	quit := time.After(time.Duration(duration) * time.Second)
	//create channel string
	var ch = make(chan string)
	//goroutine read channel
	go func() {
		for val := range ch {
			fmt.Println("Read ", val)
		}
	}()
	//infinity cycle input random string in channel, but after exit time, programm end
	for {
		select {
		case <-quit:
			close(ch)
			fmt.Println("Program end")
			os.Exit(0)
		default:
			str := RandomString(5)
			fmt.Println("Write ", str)
			ch <- str
		}
	}

}

//create random string with given length
func RandomString(cntSymb int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, cntSymb)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
