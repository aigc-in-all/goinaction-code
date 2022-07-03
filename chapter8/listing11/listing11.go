// 这个程序展示如何创建定制的日志记录器
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	// New创建一个新的Logger，out参数设置日志数据将被写入的目的地
	// 参数prefix会在生成的每行日志的最开始出现
	// 参数flag定义日志记录包含哪些属性
	/*
		func New(out io.Writer, prefix string, flag int) *Logger {
			return &Logger(out: out, perfix: prefix, flag: flag)
		}
	*/
	// Discard是一个io.Writer，所有的Write都不会有动作，但是会成功返回
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know now")
	Error.Println("Something has failed")
}

/*
INFO: 2022/07/03 13:18:21 listing11.go:41: Special Information
WARNING: 2022/07/03 13:18:21 listing11.go:42: There is something you need to know now
ERROR: 2022/07/03 13:18:21 listing11.go:43: Something has failed
*/
/*
并且ERROR会写到文件errors.txt
*/
