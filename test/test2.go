package main

import "fmt"

type P struct {
	name string
}

type N1 struct {
	*P
	age int
}

type N2 struct {
	P
	age int
}

func main() {
	m := P{"aa"}
	n1 := N1{&m, 12}
	fmt.Println(*(n1.P))

	n2 := N2{m, 13}
	fmt.Println(n2.P)

	var a []*int
	a = nil
	fmt.Println(len(a))
}
