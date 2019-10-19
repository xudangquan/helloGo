package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 8}
	fmt.Println(q)
	q[0] = 100
	fmt.Println(q)
	p := &q
	fmt.Println(p, (*p)[1])
	q[5] = 200
	fmt.Println(q)
}
