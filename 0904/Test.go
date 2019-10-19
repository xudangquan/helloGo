package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//var a int = 8
	var b int64 = 8
	s := "525"
	c := b + 2
	//if a > b {
	//	fmt.Println("=")
	//} else {
	//	fmt.Println("!=")
	//}
	fmt.Printf("%T", c)
	fmt.Println()
	int64s, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		fmt.Println(int64s)
		fmt.Printf("%T", int64s)
	}
	fmt.Println()
	fmt.Println(time.Now().Unix())
	fmt.Println("-------")
	timeStr := time.Now().Format("2006-01-02")
	parse, err := time.Parse("2006-01-02", timeStr)
	if err == nil {
		fmt.Println(timeStr, parse.Unix())
		fmt.Printf("%T", parse.Unix())
		fmt.Println()
		fmt.Println(time.Now().Unix() - parse.Unix())
	}

	fmt.Println("--------------")
	ss := strconv.FormatInt(time.Now().Unix()-parse.Unix(), 10)
	fmt.Println(ss)
	len := len(ss)
	if len < 8 {
		for i := 0; i < 8-len; i++ {
			ss = "0" + ss
		}
	}
	fmt.Println(ss)
}
