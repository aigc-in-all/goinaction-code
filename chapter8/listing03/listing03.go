// 这个示例程序展示如何使用最基本的log包
package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
func main() {
	// Println写到标准日志记录器
	log.Println("message")

	// Fatalln在调用Println()之后会接着调用os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln在调用Println()之后接着调用panic()
	log.Panicln("panic message")
}

/*Output
TRACE: 2022/06/26 23:39:48.335539 /Users/heqingbao/workspace/go_projects/goinaction-code/chapter8/listing03.go:12: message
TRACE: 2022/06/26 23:39:48.335793 /Users/heqingbao/workspace/go_projects/goinaction-code/chapter8/listing03.go:15: fatal message
exit status 1
*/
