package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
%d			十进制整数
%x,%o,%b  	十六进制，八进制，二进制整数。
%f,%g,%e 	 浮点数：	3.141593	3.141592653589793	3.141593e+00
%t			布尔：true或false
%c			字符（rune）	(Unicode码点)
%s			字符串
%q			带双引号的字符串"abc"或带单引号的字符'c'
%v			变量的自然形式（natural	format）
%T			变量的类型
%%			字面上的百分号标志（无操作数）
制表符	\t	和换行符	\n

*/
func main() {
	fmt.Println("test print")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) //默认情况下，	Printf	不会换行。按照惯例，以字 母	f	结尾的格式化函数，如	log.Printf	和	fmt.Errorf
			//以	ln	结尾的格式化函数，则遵循	Println	的方式，以跟	%v	差不多的方式格式化参数，并 在最后添加一个换行符
		}
	}
}
