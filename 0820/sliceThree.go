package main

import "fmt"

func main() {
	s := [][][]string{
		[][]string{
			[]string{
				"000", "001", "002",
			},
			[]string{
				"010",
			},
			[]string{
				"020", "021",
			},
		},
		[][]string{
			[]string{
				"100", "101", "102", "103",
			},
		},
	}
	fmt.Println(s)
	for i, _ := range s {
		for i1, _ := range s[i] {
			//fmt.Println(i, i1, s[i][i1])
			for i2, _ := range s[i][i1] {
				fmt.Println(i, i1, i2, s[i][i1][i2])
			}
		}
	}
}
