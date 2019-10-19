package main

import "fmt"

func main() {
	var arr []string
	fmt.Println(len(arr))
	if arr == nil {
		fmt.Println("this is null")
	}

	if len(arr) > 0 {
		fmt.Println("len arr > 0")
	} else {
		fmt.Println("len this is null")
	}

}
