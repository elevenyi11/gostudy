package main

import "fmt"

func main() {
	slice := make([]int, 5)
	slice = make([]int, 5, 10)

	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice ", slice)

	arr := [7]int{1, 2, 3, 4, 5, 7, 8}
	slice1 := arr[3:5]
	fmt.Println("slice1:", slice1)

	sliceA := []int{1, 2, 3, 4, 5, 6}
	sliceB := sliceA[2:4]
	fmt.Println("sliceB:", sliceB)
	sliceB = append(sliceB, 7, 8, 9, 0)
	fmt.Println("sliceB:", sliceB)
	fmt.Println("cap:", cap(sliceB))
	fmt.Println("len:", len(sliceB))

	slice2 := []int{1, 2, 3, 4, 5}
	newslice := slice2[0:3]
	fmt.Println("newslice:", newslice)
	newslice[0] = 9
	fmt.Println("slice2:", slice2)
	fmt.Println("newslice:", newslice)

	newslice1 := make([]int, len(slice2))
	copy(newslice1, slice2)
	fmt.Println("newslice1:", newslice1)
	newslice1[0] = 99
	fmt.Println("newslice1:", newslice1)
	fmt.Println("slice2:", slice2)

	copyslice := make([]int, 3)
	copy(copyslice, slice2[0:3])
	fmt.Println("copyslice", copyslice)
}
