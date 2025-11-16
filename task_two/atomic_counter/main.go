package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用 sync/atomic 实现无锁计数器的示例
// 启动 10 个 goroutine，每个执行 1000 次递增操作
func main() {
	var counter uint64
	var wg sync.WaitGroup

	goroutines := 10
	incsPerG := 1000

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incsPerG; j++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter (atomic): %d\n", counter)
}
