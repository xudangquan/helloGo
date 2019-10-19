package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan string)
	go func() {
		//fmt.Println("前奏")
		ch1 <- 1
		fmt.Println("赋值")
	}()
	fmt.Println("first")
	<-ch1
	fmt.Println("stop")
	time.Sleep(time.Second)
	//ch2 <- "s"
	//select {
	//case e1 := <-ch1:
	//	fmt.Println(e1)
	//	fmt.Printf("1th case is selected. e1=%v.\n", e1)
	//case e2 := <-ch2:
	//	fmt.Printf("1th case is selected. e2=%v.\n", e2)
	//default:
	//	fmt.Println("No data!")
	//}
}
