package gotest

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("3333")
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func TestSliceAppend(t *testing.T) {
	var s []int
	s = append(s, 1, 2)
	PrintSlice(s)
}
