package main

import "fmt"

func main() {
	var arr1, arr2 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i
	}
	printHelper("arr1", arr1)

	arr2 = [5]int{2, 3, 4, 5, 6}

	printHelper("arr2", arr2)

	var arr3 = [...]int{0, 1, 2, 4, 5}
	printHelper("arr3", arr3)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}

func printHelper(name string, arr [5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v[%v]:%v\n", name, i, arr[i])
	}

	fmt.Printf("len of %v: %v\n", name, len(arr))

	fmt.Printf("cap of %v:%v\n", name, cap(arr))
}
