package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 计数器加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待goroutine结束
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示5000以内的素数
func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

/*Output
Create Goroutines
Waiting To Finish
B: 2
B: 3
...
B: 4583
B: 4591
A: 3
A: 5
...
A: 4561
A: 4567
B: 4603
B: 4621
...
Completed B
A: 4457
A: 4463
...
A: 4993
A: 4999
Completed A
Terminating Program
*/

/*
只有一个逻辑处理器时，goroutine执行耗时操作时，会被调度到交替执行
*/
