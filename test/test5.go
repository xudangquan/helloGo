package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, 世界")
	var s string
	//s = "333,"
	strings.TrimRight(s, "0")
	fmt.Println(s)
	s = strings.TrimRight(s, "0")
	fmt.Println(s)

}
