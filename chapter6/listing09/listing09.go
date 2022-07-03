package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter int
)

func main() {

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待goroutine结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加counter变量的值
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 让当前goroutine从当前线程退出，并放回到队列等待再次被调度
		runtime.Gosched()

		value++
		counter = value
	}
}

/*Output
Final Counter: 2
*/

/*
多个goroutine在没有互相同步的情况下，访问某个共享的资源，并试图读和写这个资源，就处于相互竞争的状态，这种情况被称作竞争状态（race candition）

Go语言提供了一个工具用于检测竞争状态
$ go build -race listing09.go
$ ./listing09
==================
WARNING: DATA RACE
Read at 0x0001049f58c0 by goroutine 8:
  main.incCounter()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:33 +0x80
  main.main.func2()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:21 +0x30

Previous write at 0x0001049f58c0 by goroutine 7:
  main.incCounter()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:39 +0xa0
  main.main.func1()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:20 +0x30

Goroutine 8 (running) created at:
  main.main()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:21 +0x48

Goroutine 7 (finished) created at:
  main.main()
      /Users/heqingbao/workspace/go_projects/goinaction-code/chapter6/listing09/listing09.go:20 +0x3c
==================
Final Counter: 4
Found 1 data race(s)
*/
