package main

import "fmt"

func main() {
	var err error
	if err != nil {
		fmt.Println("no nil")
	} else {
		fmt.Println("nil")
	}
}
