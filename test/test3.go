package main

import (
	"fmt"
	"reflect"
)

type OO struct {
	name string
}

func main() {
	a := OO{}
	if a.name == "1" {
		fmt.Println("111")
	}
	fmt.Println(a)
	fmt.Println("222")

	v1 := "123456"
	v2 := 12

	// reflect
	fmt.Println("v1 type:", reflect.TypeOf(v1))
	fmt.Println("v2 type:", reflect.TypeOf(v2))
	if reflect.TypeOf(v1) == reflect.TypeOf("") {
		fmt.Println("555")
	}
}
