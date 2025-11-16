package main

import (
	"fmt"
	"sync"
	"time"
)

// 示例1：无缓冲通道，生产者发送 1..10，消费者接收并打印
func demoUnbuffered() {
	fmt.Println("示例1：无缓冲通道，两个协程通信（1..10）")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
		close(ch)
	}()

	// 消费者
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收:", v)
		}
	}()

	wg.Wait()
	fmt.Println("示例1 完成\n")
}

// 示例2：带缓冲通道，生产者发送 100 个整数，消费者接收并打印
func demoBuffered() {
	fmt.Println("示例2：带缓冲通道，生产者发送100个整数")
	// 缓冲大小设为 20，示范缓冲机制
	ch := make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收:", v)
		}
	}()

	wg.Wait()
	fmt.Println("示例2 完成")
}

func main() {
	demoUnbuffered()
	demoBuffered()
}
