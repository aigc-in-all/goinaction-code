// 这个示例程序展示如何使用互拆锁来定义一段需要同步访问的代码临界区资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutext  sync.Mutex
)

func main() {

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter 使用互拆锁来同步并保存安全访问
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutext.Lock()
		{
			value := counter

			// 当前goroutine从线程退出，并放到队列
			runtime.Gosched()

			value++

			counter = value
		}
		// 释放锁，允许其它正在等待的goroutine进入临界区
		mutext.Unlock()
	}
}
