package asynchronous

import (
	"fmt"
	"time"
	"./async"
)

func DoAsync() int {
	fmt.Println("Warming up ...")

	time.Sleep(3 * time.Second)

	fmt.Println("Done ...")

	return 1
}

func DoTest() {
	fmt.Println("Let's start ...")
	future := async.Exec(func() interface{} {
		return DoAsync()
	})

	fmt.Println("Done ...")

	result := future.Await()
	fmt.Println(result)
}
