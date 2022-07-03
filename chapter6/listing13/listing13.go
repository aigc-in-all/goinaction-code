// 这个示例程序展示如何使用atomic包来提供对数值类型的安全访问
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	wg      sync.WaitGroup
	counter int64
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
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1)

		// 让当前goroutine从当前线程退出，并放回到队列等待再次被调度
		runtime.Gosched()
	}
}

/*Output
Final Counter: 4
*/
