package main

import (
	"fmt"
)

var recursionMap = make(map[int]int, 50)

func main() {
	array := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	fmt.Println(array)
	InsertionSort2(array)
	fmt.Println(array)

	fmt.Printf("recursion(10)=%v", recursion(10))
	fmt.Printf("recursion(999)=%v", recursion(999))
	a := []int{4, 5, 6, 3, 2, 1}
	aResult := bubbleSort(a)
	fmt.Println("bubbleSort(a)=", aResult)

	fmt.Println("4 >> 1 =", 4>>1)

	fmt.Println("4 << 1 =", 4<<3)
}

func stackTest() {
	/*
		var str = []string{"(", "5", "+", "6", ")", "*", "3"}
		fmt.Println(reflect.TypeOf(""))
		var stack = NewStack(reflect.TypeOf(""))
		for _, s := range str {
			if s == "(" {
				stack.Push(s)
			} else if s == ")" {
				stack.Pop()
			}
		}
		fmt.Println(stack.Len())
	*/
}

// 递归
func recursion(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	if value, ok := recursionMap[n]; ok {
		return value
	} else {
		result := recursion(n-1) + recursion(n-2)
		recursionMap[n] = result
		return result
	}
}

// 冒泡排序
func bubbleSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return a
}

// 插入排序
func InsertionSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		// 查找插入位置
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j] //数据移动
			} else {
				break
			}
		}
		a[j+1] = value //插入数据
	}
	return a
}

//插入排序2
//https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F#Go
func InsertionSort2(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			} else {
				break
			}
		}
	}
}

// 二分查找
func bsearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for {
		if low <= high {
			mid := (low + high) / 2
			if a[mid] == value {
				return mid
			} else if a[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			return -1
		}
	}
}

// 二分查找递归实现
func bsearch2(a []int, n int, val int) int {
	return bsearchInternally(a, 0, n-1, val)
}

func bsearchInternally(a []int, low int, high int, val int) int {
	if low > high {
		return -1
	}

	mid := low + ((high - low) >> 1) //相比除法运算来说，计算机处理位运算要快得多
	if a[mid] == val {
		return mid
	} else if a[mid] < val {
		return bsearchInternally(a, mid+1, high, val)
	} else {
		return bsearchInternally(a, low, mid-1, val)
	}
}
