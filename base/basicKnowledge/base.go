package basicKnowledge

import (
	//这个语法引入了 encoding/base64 包并使用名称 b64代替默认的 base64。这样可以节省点空间。
	b64 "encoding/base64"
	"fmt"
	"time"
)

/*
%d 十进制整数
%x,%o,%b 十六进制，八进制，二进制整数
%f,%g,%e 浮点数：3.141593,3.141592653589793， 3.141593e+00
%t 布尔：true或false
%c 字符（rune）（Unicode码点）
%s 字符串
%q 带双引号的字符串"abc"或带单引号的字符'c'
%v 变量的自然形式(natural format)
%T 变量的类型
%% 字面上的百分号标识（无操作数）
*/


func Base(){
	fmt.Print(1)
	fmt.Println(fmt.Sprintf("TestBase:%v ", time.Now().Format("2006-01-02 15:04:05")))
}

func Base64(){
	//标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。
	data := "abc123!?$*&()'-=@~"
	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(fmt.Sprintf("sEnc:%v ", sEnc))

	//解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(fmt.Sprintf("sDec:%s ", string(sDec)))

	//使用 URL 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

func ForTest(){
	i :=1
	for i <3 {
		fmt.Println(i)
		i +=1
	}
}

