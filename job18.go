package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	 По завершению программа должна выводить итоговое значение счетчика.
*/

/* CASE 1 */

// //create struct with mutex
// type Counter_t struct {
// 	cnt uint
// 	mtx sync.Mutex
// }

// //function increment with mutex
// func (c *Counter_t) Inc() {
// 	c.mtx.Lock()
// 	c.cnt++
// 	c.mtx.Unlock()
// }

// //function get counter
// func (c *Counter_t) GetIncCounter() uint {
// 	return c.cnt
// }

// func main() {
// 	//create struct
// 	s := Counter_t{}
// 	//create synchronization primitive
// 	var wg sync.WaitGroup
// 	//goroutine increment struct 10000 times
// 	for i := 0; i < 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			s.Inc()
// 			wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// 	//output increment
// 	fmt.Println(s.GetIncCounter())

// }

/* CASE 2 */

//counter struct
type Counter_t struct {
	cnt uint64
}

//func increment atomic package
func (c *Counter_t) Inc() {
	atomic.AddUint64(&c.cnt, 1)
}

//func get counter
func (c *Counter_t) GetIncCounter() uint64 {
	return c.cnt
}

func main() {
	//create struct
	s := Counter_t{}
	//create synchronization primitive
	var wg sync.WaitGroup
	//goroutine increment struct 10000 times
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			s.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	//output increment
	fmt.Println(s.GetIncCounter())
}
