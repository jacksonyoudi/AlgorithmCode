package main

import (
	"fmt"
	"sync"
)

type M struct {
	m      map[int]int
	locker *sync.Mutex
}

func (t *M) Set(k, v int) {
	t.locker.Lock()
	defer t.locker.Unlock()
	fmt.Println("done", k)
	t.m[k] = v
}
func (t M) Get(k int) int {
	fmt.Printf("%p\n", &t.locker)
	t.locker.Lock()
	defer t.locker.Unlock()
	fmt.Println("done", k)
	return t.m[k]
}


type A struct{
	m      map[int]int
	locker *sync.Mutex
}



func main() {
	m := M{
	}
	var wg sync.WaitGroup
	m.m = make(map[int]int, 1000)
	m.locker = new(sync.Mutex)

	wg.Add(100)
	for i := 0; i < 100; i++ {

		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				m.Set(i, i)
			} else {
				m.Get(i)
			}
		}(i)
	}
	wg.Wait()


}
