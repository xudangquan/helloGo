package main

import (
	"fmt"
)

type M struct {
	name string
}

func main() {
	//var b *M
	//a := M{"zhang"}
	//b = &a
	//fmt.Println(a.name)
	//fmt.Println(b.name)
	var c []*M
	if false {
		c = []*M{{"n"}, {"m"}}
	}
	//fmt.Println(c)
	//fmt.Println(reflect.TypeOf(c))
	if len(c) == 0 {
		fmt.Println("==空")
	} else {
		fmt.Println("!=空")
		fmt.Println(len(c))
	}
	i := 8
	var aa interface{}
	if i > 0 {
		aa = &M{"a"}
	} else {
		aa = 8
	}
	fmt.Println(aa)

	var bb []*M
	bb = []*M{{"a"}, {"b"}}
	fmt.Println(bb[2].name)
}
