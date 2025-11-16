package main

import (
	"fmt"
	"sync"
)

// 使用 sync.Mutex 保护共享计数器的示例
// 启动 10 个 goroutine，每个执行 1000 次递增操作
func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	goroutines := 10
	incsPerG := 1000

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incsPerG; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter (mutex): %d\n", counter)
}
