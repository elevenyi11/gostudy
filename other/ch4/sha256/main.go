package main

import "crypto/sha256"
import "fmt"

func main() {
	//需要注意Printf函数的%x副词参数，它用于指定以十六进制的格式打印数组或 slice全部的元素，%t副词参数是用于打印布尔型数据，%T副词参数是用于显示一个值对应的 数据类型。

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
