package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	nowTime := time.Now()
	dayStr := nowTime.Format("20060102")
	// 生成当天时间的秒数
	//timeStr := time.Now().Format("2006-01-02")
	parse, err := time.Parse("20060102", dayStr)
	if err != nil {
		return // TODO
	}
	nowSecond := nowTime.Unix() - parse.Unix()
	nowTimeStamp := strconv.FormatInt(nowSecond, 10)
	fmt.Println(dayStr, nowSecond, nowTimeStamp)
	fmt.Printf("%T", nowTimeStamp)
	fmt.Printf("%T", nowSecond)
	fmt.Printf("%T", dayStr)
	fmt.Println()
	maxSeq := 9999999
	fmt.Println(strconv.Itoa(maxSeq))
	fmt.Printf("%T", maxSeq)
	fmt.Println()
	fmt.Println(time.Now().Format("20060102150405Z0700"))
}
