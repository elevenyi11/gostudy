package basicKnowledge

import (
	//这个语法引入了 encoding/base64 包并使用名称 b64代替默认的 base64。这样可以节省点空间。
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
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
基本类型列表
​类型        长度     说明
bool         1      true/false,默认false, 不能把非0值当做true(不用数字代表true/false)
byte         1      uint8 别名
rune         4      int32别名。 代表一个unicode code point
int/unit            所运行的平台，32bit/64bit
int8/uint8   1     -128 ~ 127; 0 ~ 255
int16/uint16 2     -32768 ~ 32767; 0 ~ 65535
int32/uint32 4     -21亿 ~ 21亿， 0 ~ 42亿
int64/uint64 8

float32      4     精确到7位小数,相当于c的float
float64      8     精确到15位小数,相当于c的double
complex64    8
complex128   16
uintptr            足够保存指针的32位、64位整数,指针(可以存指针的整数型)
array              值类型,数组
struct             值类型,结构体
string             值类型,字符串类型，常用
slice              引用类型,切片
map                引用类型,字典
channel            引用类型,通道
interface          接口类型,接口
function           函数类型,函数

运算符全部是从左到右结合的
优先级    运算符                        说明
  高   * / % << >> & &^(AND NOT)
       + - ! ^
       == != < <= > >=
       <-                             channel运算符
       &&
  低   ||
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

func TestBit(){
	var n uint8 = 6
	fmt.Printf("%s %08b %08b %v\n\n", "6<<1 左移1位", n, n<<1, n<<1)
	fmt.Printf("%s %08b %08b %v\n\n", "6>>1 右移1位", n, n>>1, n>>1)
	fmt.Printf("%s %08b %08b %v\n\n", "^6 位取反", n, ^n, ^n)
	fmt.Printf("%s %08b %08b %08b %v\n\n", "6&5 位与", n, 5, n&5, n&5)
	fmt.Printf("%s %08b %08b %08b %v\n\n", "6|5 位或", n, 5, n|5, n|5)
	fmt.Printf("%s %08b %08b %08b %v\n\n", "6^5 位异", n, 5, n^5, n^5)
}

func TestTime(){
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
}

func TestOsArgs(){
	// os.Args方式
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("Hello 世界!")
	} else {
		fmt.Println("Hello ", args[1]) // 第二个参数，第一个参数为命令名
	}

	// flag.Args方式
	flag.Parse()
	var ch []string = flag.Args()
	if ch != nil && len(ch) > 0 {
		fmt.Println("Hello ", ch[0]) // 第一个参数开始
	}
}