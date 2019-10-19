package main

import (
	"fmt"
	"time"
)

func main() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	//for data := range ch {
	//	fmt.Println(data)
	//	if data == 0 {
	//		break
	//	}
	//}
	v1, v2, v3, v4 := <-ch, <-ch, <-ch, <-ch
	fmt.Println(v1, v2, v3, v4)
}
