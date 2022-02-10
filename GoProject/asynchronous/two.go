package asynchronous

import (
	"fmt"
	"time"
)

func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("staring....")

	// 放到一个协程里
	go func() {
		time.Sleep(time.Second * 10)
		r <- 1
		fmt.Println("Done")
	}()

	return r

}
