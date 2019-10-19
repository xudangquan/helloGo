package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	PrintSlice(s)
	// 切片的切片
	board := [][]string{
		[]string{
			"A", "B",
		},
		[]string{
			"D", "E", "F",
		},
		[]string{
			"G", "H", "I",
		},
	}
	for i, _ := range board {
		for i1, _ := range board[i] {
			fmt.Println(i, i1, board[i][i1])
		}
	}
	fmt.Println(board)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
