package basicKnowledge

import "strconv"

//golang strconv.ParseInt 是将字符串转换为数字的函数,功能灰常之强大.
//参数1 数字的字符串形式
//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
//func ParseInt(s string, base int, bitSize int) (i int64, err error)
func TestConv(){
	i, err := strconv.ParseInt("123",10,32)
	if err != nil{
		panic(err)
	}
	println(i)
}
