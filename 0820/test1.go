package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for i, v := range a { //i,v从a复制的对象里提取出
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a) //输出[1 200 300]
		}
		a[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
	}
	fmt.Println(a) //输出[101, 102, 103]

	b := []int{1, 2, 3} //改成slice
	for i, v := range b {
		if i == 0 {
			b[1], b[2] = 200, 300
			fmt.Println(b) //输出[1 200 300]
		}
		b[i] = v + 100 //v是引用对象里的元素[1, 200, 300]
	}
	fmt.Println(b) //输出[101 300 400]
}
