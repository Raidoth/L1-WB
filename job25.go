package main

import (
	"fmt"
	"time"
)

/*

25.	Реализовать собственную функцию sleep.
*/

/* CASE 1 */

//function wait duration and return

// func Sleep(duration time.Duration) {
// 	<-time.After(duration)
// }

/* CASE 2 */

// func Sleep(duration time.Duration) {
// 	//create timer with duration
// 	t := time.NewTimer(duration)
// 	//after time return
// 	<-t.C
// }

/* CASE 3 */

func Sleep(duration time.Duration) {
	//t equal now time and duration
	t := time.Duration(time.Now().Add(duration).Unix())
	//while t > now time compute cycle, after return
	for t > time.Duration(time.Now().Unix()) {

	}
}

func main() {

	fmt.Println("Wait sleep")
	Sleep(2 * time.Second)
	fmt.Println("Sleep is done")

}
