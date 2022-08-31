package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
6.	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	/* CASE 1 */

	//create channel and start goroutine, if channel close, then goroutine end
	quit := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Goroutine end with close channel in select")
				wg.Done()
				return
			}
		}
	}()
	close(quit)
	wg.Wait()
	/* CASE 2 */

	//create channel and start goroutine, if channel close, then goroutine end
	qt := make(chan int)
	wg.Add(1)
	go func() {
		defer fmt.Println("Goroutine end with close channel in if-else")
		for {
			if val, ok := <-qt; ok {
				fmt.Println(val)
			}
			wg.Done()
			return
		}
	}()
	close(qt)
	wg.Wait()
	/* CASE 3 */

	//goroutine end with function Goexit
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Goroutine end runtime")
		for {
			time.Sleep(time.Second)
			runtime.Goexit()
		}
	}()
	wg.Wait()
	/* CASE 4 */

	//create context with cancel, if call cancel then signal send <-ctx.Done()
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine end with context")
				return
			}
		}
	}(ctx)
	cancel()
	wg.Wait()
	/* CASE 5 */

	//goroutine end with main function, but no run defer
	go func() {
		defer fmt.Println("End")
	}()
}
