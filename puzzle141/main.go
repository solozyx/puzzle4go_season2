package main

import "fmt"

//写出下面代码输出内容。

func main() {
	defer_call()
}
func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发异常")
}

//考点：defer执行顺序
//
//解答： defer 是后进先出。
//panic 需要等defer 结束后才会向上传递。出现panic恐慌时候，
//会先按照defer的后入先出的顺序执行，最后才会执行panic。

//打印后
//打印中
//打印前
//panic: 触发异常

//C:\Users\Administrator\AppData\Local\Temp\___go_build_main_go.exe #gosetup
//打印后
//打印中
//打印前
//panic: 触发异常
//
//goroutine 1 [running]:
//main.defer_call()
//        D:/dev/interview/puzzle4go_season2/puzzle141/main.go:20 +0x85
//main.main()
//        D:/dev/interview/puzzle4go_season2/puzzle141/main.go:8 +0x27
//
//Process finished with exit code 2
