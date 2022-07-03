// 这个示例程序展示如何用无缓冲的通道来模拟4个goroutine间的接力比赛”
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者将计数加1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	runner := <-baton

	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Batln\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

/*Outout
Runner 1 Running With Batln
Runner 2 To The Line
Runner 1 Exchange With Runner 2
Runner 2 Running With Batln
Runner 3 To The Line
Runner 2 Exchange With Runner 3
Runner 3 Running With Batln
Runner 4 To The Line
Runner 3 Exchange With Runner 4
Runner 4 Running With Batln
Runner 4 Finished, Race Over
*/
