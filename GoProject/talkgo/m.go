package talkgo

import (
	"context"
	"fmt"
	"os"
)

/**
 * 编写程序：并发调用 API
 * 功能：
 *   1. 有若干个API需要调用：
 *     1.1 互相之间没有任何依赖关系（即不需要保证顺序）
 *     1.2 各个API的调用耗时互不相同，也无法提前预知
 *   2. 要求：
 *     2.1 要求尽快完成调用，并返回结果
 *     2.2 最大开 N个 并发
 *     2.3 中间任意一个调用失败，剩余未处理的调用全部取消，并立即返回错误
 */

const N = 3 // 最大并发数

func main() {
	var apiList = []string{"a", "b", "c", "d", "e", "f", "g"} // 待调用API列表（API个数会动态变化，非固定）
	controlApi(apiList)
}

func controlApi(apiList []string) {
	dataCh := make(chan string, N)
	end := make(chan bool)
	ctx, _ := context.WithCancel(context.Background())

	go production(dataCh, apiList, end)
	go consume(ctx, dataCh, end)

	for {
		select {
		case data := <-end:
			fmt.Println(data)


		case <-ctx.Done():
			fmt.Println("end------------------")
			os.Exit(1)
		}
	}
}

func callApi(ctx context.Context, api string) error {
	// handling api here
	return nil
}

// 生产者
func production(ch chan string, apiList []string, end chan bool) {
	for _, i := range apiList {
		ch <- i
	}

	end <- true
}

// 消费者
func consume(ctx context.Context, ch chan string, end chan bool) {
	for {
		select {
		case data := <-ch:
			err := callApi(ctx, data)
			if err != nil {
				ctx.Done()
			}

			fmt.Println(data)
		case <-ctx.Done():
			fmt.Println("end------------------")
		}
	}
	end <- true
}
