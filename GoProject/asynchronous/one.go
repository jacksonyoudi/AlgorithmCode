package asynchronous

import (
	"fmt"
	"sync"
)

// 异步将整数列表发送到通道，关闭通道返回；
// gen函数,nums表示多个int类型数据，多参数  返回一个只读取的通道, 不会阻塞
func gen(nums ...int) <-chan int {
	out := make(chan int)

	// 用一个协程 在后台运行
	go func() {
		for _, n := range nums {
			out <- n
		}
		// 关闭， 只能读，不能写了
		// close 后不能写入,其他goroutine读取out的数据完之后就会退出不在阻塞
		close(out)
	}()

	return out
}

// 将通道数据读取，异步返回保存读取数据的平方数通道
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	//异步的从in获取数据，开始是阻塞的，遍历写入out，当gen调用close(out）后就会不在阻塞等待in的数据
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	// 输出通道
	out := make(chan int)

	outPut := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// cs 数量的信号量
	wg.Add(len(cs))

	for _, c := range cs {
		go outPut(c)
	}

	// 等待输入通道所有数据全部写入输出通道
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func test() {
	// 返回都是 channel
	c := gen(2, 3)
	out := sq(c)
	out1 := sq(c)

	// 使用close 就会
	for i := range merge(out, out1) {
		fmt.Println(i)
	}
}
