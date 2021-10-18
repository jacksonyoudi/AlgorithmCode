package tx

import (
	"context"
	"os"
	"sync"
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
	cch := make(chan string, N)

	//生产者
	for _, i := range apiList {
		cch <- i
	}

	contextMap := sync.Map{}
	errChan := make(chan error)

	for j := 0; j < N; j++ {
		go func(cch chan string) {

			for {
				m := <-cch
				background := context.Background()
				contextMap.Store(m, &background)
				err := callApi(background, m)
				if err != nil {
					errChan <- err
				}
			}
		}(cch)
	}

	for {
		<-errChan

		contextMap.Range( func(key, value interface{}) bool {
			value.(context.Context).Done()

			return false
		})


		os.Exit(1)
	}

}

func callApi(ctx context.Context, api string) error {
	// handling api here
	return nil
}
