package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Printf("%T %T %T\n", months, Q2, summer)
	fmt.Println(Q2)
	fmt.Println(summer)
	popSameMonth(Q2, summer)

	a := [...]int{0, 1, 2, 3, 4, 5} //array
	reverse(a[:])
	fmt.Println(a)
	s := []int{0, 1, 2, 3, 4, 5} //silce slice和数组的字面值 语法很类似，它们都是用花括弧包含一系列的初始化元素，但是对于slice并没有指明序列的 长度。这会隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组。就像数组字126S lice面值一样，slice的字面值也可以按顺序指定初始化值序列，或者是通过索引和元素值指定， 或者的两种风格的混合语法初始化。slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有 全部相等元素。不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相 等（[]byte），但是对于其他类型的slice，我们必须自己展开每个元素进行比较.
	reverse(s[:2])
	fmt.Println(s)

	reverse(s[2:])
	fmt.Println(s)

	reverse(s)
	fmt.Println(s)
}

func popSameMonth(months1 []string, months2 []string) {
	for _, s := range months1 {
		for _, q := range months2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
