package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1
	x := 3.14
	y := "hello"
	fmt.Println(reflect.TypeOf(x))
	fmt.Printf("%T\n", y)

	// 2
	ifa := make([]interface{}, 3)
	ifa[0] = 1
	ifa[1] = "hello"
	for _, data := range ifa {
		if value, ok := data.(string); ok == true {
			fmt.Println(value)
		} else if value, ok := data.(int); ok == true {
			fmt.Println(value)
		}
	}

	// 3
	for _, data := range ifa {
		switch value := data.(type) {
		case int:
			fmt.Printf("int型 %d\n", value)
		case string:
			fmt.Println("string型", value)
		}
	}
}
