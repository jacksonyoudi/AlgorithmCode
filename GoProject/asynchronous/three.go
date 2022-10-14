package asynchronous

import (
	"fmt"
	"time"
)

func DoAsync() int {
	fmt.Println("Warming up ...")

	time.Sleep(3 * time.Second)

	fmt.Println("Done ...")

	return 1
}

func DoTest() {
	fmt.Println("Let's start ...")
	future := Exec(func() interface{} {
		return DoAsync()
	})

	fmt.Println("Done ...")

	result := future.Await()
	fmt.Println(result)
}
