package basicKnowledge

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)


func stringFunc(){
	//这个函数的作用是按照1：n个空格来分割字符串最后返回的是
	//[]string的切片
	fmt.Println(strings.Fields("hello widuu golang")) //out  [hello widuu golang]

	fmt.Println(strings.FieldsFunc("widuunhellonword", split)) // [widuu hello word]根据n字符分割

	//这个函数是将一个[]string的切片通过分隔符，分割成一个字符串
	s := []string{"hello", "word", "xiaowei"}
	fmt.Println(strings.Join(s, "-")) // hello-word-xiaowei

	//Split这个就是把字符串按照指定的分隔符切割成slice
	fmt.Println(strings.Split("a,b,c,d,e", ",")) //[a b c d e]

	//在前边的切割完成之后再后边在加上sep分割符
	fmt.Println(strings.SplitAfter("a,b,c,d", ",")) //[a, b, c, d]

	//该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 4)) //["a," "b," "c," "d,r"]
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 5)) //["a," "b," "c," "d," "r"]

	//切割字符串的时候自己定义长度，如果sep为空，那么每一个字符都分割
	fmt.Println(strings.SplitN("a,b,c", ",", 2)) //[a b,c]

	//查找某个字符是否在这个字符串中存在，存在返回true
	fmt.Println(strings.Contains("widuu", "wi")) //true
	fmt.Println(strings.Contains("wi", "widuu")) //false

	//查询字符串中是否包含多个字符
	fmt.Println(strings.ContainsAny("widuu", "w&d")) //true

	//字符串中是否包含rune类型，其中rune类型是utf8.RUneCountString可以完整表示全部Unicode字符的类型
	fmt.Println(strings.ContainsRune("widuu", rune('w'))) //true
	fmt.Println(strings.ContainsRune("widuu", 20))        //fasle

	//在一段字符串中有多少匹配到的字符
	fmt.Println(strings.Count("widuu", "uu")) //1
	fmt.Println(strings.Count("widuu", "u"))  //2

	//查找字符串，然后返回当前的位置，输入的都是string类型，然后int的位置信息
	fmt.Println(strings.Index("widuu", "i")) //1
	fmt.Println(strings.Index("widuu", "u")) //3

	//字符串第一次出现的位置，如果不存在就返回-1
	fmt.Println(strings.IndexAny("widuu", "u")) //3

	//查找第一次出现的位置，只不过这次C是byte类型的，查找到返回位置，找不到返回-1
	fmt.Println(strings.IndexByte("hello xiaowei", 'x')) //6

	//查找rune类型的
	fmt.Println(strings.IndexRune("widuu", rune('w'))) //0

	//是通过类型的转换来用函数查找位置
	fmt.Println(strings.IndexFunc("nihaoma", split2)) //3

	//最后出现的位置，正好跟index相反
	fmt.Println(strings.LastIndex("widuu", "u")) // 4

	//这个跟indexAny正好相反，查找最后一个
	fmt.Println(strings.LastIndexAny("widuu", "u")) // 4
}

func split2(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}

func split(s rune) bool {
	if s == 'n' {
		return true
	}
	return false
}


func u2s(from string)(to string, err error){
	bs, err := hex.DecodeString(strings.Replace(from,`\u`,``, -1))
	if err !=nil{
		return
	}
	for i,bl,br,r := 0, len(bs),bytes.NewReader(bs), uint16(0); i <bl ; i+=2{
		binary.Read(br ,binary.BigEndian,&r)
		to += string(r)
	}
	return

}
