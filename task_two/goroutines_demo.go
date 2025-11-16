package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// printOdds 打印 1 到 10 的奇数
func printOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("odd: %d\n", i)
		time.Sleep(30 * time.Millisecond)
	}
}

// printEvens 打印 2 到 10 的偶数
func printEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("even: %d\n", i)
		time.Sleep(40 * time.Millisecond)
	}
}

// ScheduleTasks 接收一组任务并并发执行，返回每个任务的执行耗时
func ScheduleTasks(tasks []func()) []time.Duration {
	var wg sync.WaitGroup
	durations := make([]time.Duration, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			durations[i] = time.Since(start)
		}(i, task)
	}

	wg.Wait()
	return durations
}

func main() {
	fmt.Println("Demo 1: print odd and even using goroutines")
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdds(&wg)
	go printEvens(&wg)
	wg.Wait()

	fmt.Println("\nDemo 2: task scheduler measuring durations")
	rand.Seed(time.Now().UnixNano())

	tasks := []func(){
		func() {
			d := time.Duration(rand.Intn(400)+100) * time.Millisecond
			time.Sleep(d)
			fmt.Println("task 1 done")
		},
		func() {
			d := time.Duration(rand.Intn(400)+100) * time.Millisecond
			time.Sleep(d)
			fmt.Println("task 2 done")
		},
		func() {
			d := time.Duration(rand.Intn(400)+100) * time.Millisecond
			time.Sleep(d)
			fmt.Println("task 3 done")
		},
	}

	durations := ScheduleTasks(tasks)
	for i, d := range durations {
		fmt.Printf("task %d duration: %v\n", i+1, d)
	}
}
