package main

import (
	"fmt"
	"time"
)

func main() {
	//for range time.Tick(time.Second * 3){
	//	fmt.Println("hello world!")
	//}
	c := time.Tick(time.Second * 3)
	for {
		a := <-c
		go f()
		fmt.Println(a)
	}
}

func f() {
	fmt.Println("hello world")
}
