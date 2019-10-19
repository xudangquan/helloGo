package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch1 <- 1
	<-ch1
	select {
	case e1 := <-ch1:
		fmt.Println(e1)
		fmt.Printf("1th case is selected. e1=%v.\n", e1)
	case e2 := <-ch2:
		fmt.Printf("1th case is selected. e2=%v.\n", e2)
	default:
		fmt.Println("No data!")
	}
}
