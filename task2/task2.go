package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
func pointer1(num *int) {
	*num += 10
}

func pointerSlice(nums *[]int) {
	//for i := 0; i < len(*nums); i++ {
	//	(*nums)[i] *= 2
	//}
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// goroutine
func goroutine1() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("1...10的奇数： %d\n", i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("1...10的偶数： %d\n", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("goroutine end\n")
}

type Task func()

func schedule(tasks []Task) {
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		go func(task Task, id int) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务 %d 执行完成, 耗时: %v\n", id, duration)
		}(task, i)
	}
	wg.Wait()
	fmt.Printf("all schedule success\n")
}

// Channel
func bufferedChannel() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		time.Sleep(5 * time.Second)
	}()

	go func() {
		for v := range ch {
			fmt.Println("读取缓冲channel：", v)
		}
	}()

	time.Sleep(10 * time.Second)
}

// 锁机 累加计数
func addGoroCount() {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			count++
		}()
	}
	wg.Wait()
	fmt.Printf("最终计数: %d\n", count)
}

// 原子累计计数
func atomicAddCount() {
	var wg sync.WaitGroup
	var count int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println("原子操作最终计数:", count)
}

func main() {
	//num := 4
	//pointer1(&num)
	//fmt.Println(num)
	//
	//nums := []int{1, 6, 3, 4, 5}
	//pointerSlice(&nums)
	//fmt.Println(nums)

	//goroutine1()
	//tasks := []Task{
	//	func() { time.Sleep(800 * time.Microsecond) },
	//	func() { time.Sleep(5 * time.Second) },
	//	func() { time.Sleep(200 * time.Microsecond) },
	//}
	//schedule(tasks)

	//bufferedChannel()
	//addGoroCount()
	atomicAddCount()
}
