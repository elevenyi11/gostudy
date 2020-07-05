package main

import (
	"fmt"
	"testing"
)

func TestSliceInt(t *testing.T) {
	fmt.Println("int slice start")
	slice := NewSlice()
	slice.Add(1)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(3)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(3)
	fmt.Println("current int slice:", slice)
	fmt.Println("int slice end")
}
