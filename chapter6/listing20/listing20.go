// 这个示例程序展示如何使用无缓冲的通道来模拟2个goroutine间的网球比赛
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
	court := make(chan int)

	wg.Add(2)

	// 启动两个选手
	go player("A", court)
	go player("B", court)

	// 发球
	court <- 1

	wg.Wait()
}

// player 模拟一个选手在找网球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭了，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}

/*Output
Player B Hit 1
Player A Hit 2
Player B Hit 3
Player A Missed
Player B Won
*/
