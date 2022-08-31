package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
7.	Реализовать конкурентную запись данных в map
*/

//create struct map with mutex
type ThreadMap struct {
	maps map[string]string
	mtx  sync.RWMutex
}

// function create thread map
func CreateTMap() *ThreadMap {
	return &ThreadMap{
		maps: make(map[string]string, 5),
	}
}

//function add k,v in thread map with mutex
func (t *ThreadMap) AddMap(key, value string) {
	t.mtx.Lock()
	t.maps[key] = value
	t.mtx.Unlock()
}

//function show all thread map
func (t *ThreadMap) ViewAllMap() {
	t.mtx.RLock()
	fmt.Println("Map have: ")
	var cnt int
	for key, val := range t.maps {
		cnt++
		fmt.Printf("%d Key: %s, Value: %s\n", cnt, key, val)
	}
	t.mtx.RUnlock()

}

func main() {
	//create synchronization primitive
	var wg sync.WaitGroup
	//create thread map
	s := CreateTMap()
	//goroutine add in map 100 string
	go func() {
		wg.Add(1)
		for i := 0; i < 100; i++ {
			s.AddMap(RandomString(5), RandomString(2))
		}
		wg.Done()
	}()
	//goroutine add in map 50 string
	go func() {
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go s.AddMap(RandomString(3), RandomString(1))
			wg.Done()
		}

	}()
	//goroutine add in map 15 string
	go func() {
		wg.Add(1)
		for i := 0; i < 15; i++ {
			s.AddMap(RandomString(8), RandomString(3))
		}
		wg.Done()
	}()
	//goroutine show map 3 times
	go func() {
		wg.Add(1)
		for i := 0; i < 3; i++ {
			fmt.Println("View map: ", i)
			s.ViewAllMap()
		}
		wg.Done()
	}()
	wg.Wait()
	//show full map, after all goroutines
	s.ViewAllMap()

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
