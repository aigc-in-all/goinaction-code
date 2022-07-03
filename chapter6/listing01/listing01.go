package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg用来等待程序完成，计数器加2，表示要等待两个goroutine执行完成
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待goroutine执行结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

/*Output
Start Goroutines
Waiting To Finish
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z
Terminating Program
*/

/*
当只有一个逻辑处理器时，gorouting是顺序执行（因为执行太快了未发生调度），并不会有并发的效果
*/
