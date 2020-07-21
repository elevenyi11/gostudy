package basicKnowledge

import (
	"bytes"
	"fmt"
	"testing"
)

//字符串底层就是一个字节数组
//当你创建一个字符串时，其本质就是一个字节的数组。这意味着你可以像访问数组一样的访问单独的某个字节。
func TestString1(t *testing.T){
	str := "hello"
	for i := 0; i < len(str) ; i++ {
		fmt.Printf("%b %s\n", str[i], string(str[i]))
	}
}

func TestString2(t *testing.T){
	str := "something"
	buf := bytes.NewBufferString(str)
	for i := 0; i < 1000; i++ {
		buf.Write([]byte(randomString()))
	}
	fmt.Println(buf.String())
}
func randomString() string{
	ret :="pretend-this-is-random"
	return ret
}

// 使用字节数组会进一步提升上述代码的效率，但你需要知道最终字符串的大小。一个直观的例子就是 Go 语言中的 left-pad 实现。
// 你可以像拼接其他数组一样拼接字符串
//当你需要截取字符串中的一部分时，可以使用像截取数组某部分那样来操作
//示例代码：
func TestString3(t *testing.T){
	str :="XBodycontentX"
	content := str[1:len(str) -1]
	fmt.Println(content)
}
//使用 ` 符号来创建多行字符串
//你希望在代码中定义一个包含多行地址信息的字符串，那么你需要用到 ` 这个字符，
//如下所示：
func TestString4(t *testing.T){
	str := `Mr. Smith
123 Something St 
Some City, CA 94043`
	fmt.Println(str)
}

//你可以在字符串中嵌入 Unicode 字符
//假设实现 WebSocket 通讯时，你需要让传输的数据以 0x00 开始，以 0xFF 结束[源码]
func TestString5(t *testing.T){
	str := "\x00BodyContent\xff"
	fmt.Println(str)
//同样的，你可以使用 Unicode 字符串来处理，或者也可以在字符串中使用原始字符。例如，下面的代码都是有效的：
	a := "ÿay!"
	b := "\u00FFay!"
	fmt.Println(a, b)
}

func TestU2s(t *testing.T){
	str :=`\u5bb6\u65cf`
	fmt.Println(u2s(str))
}
