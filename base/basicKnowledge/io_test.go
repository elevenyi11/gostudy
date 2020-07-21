package basicKnowledge

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReadDir(t *testing.T){
	dir_list, e := ioutil.ReadDir("d:/test")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
		fmt.Println(v.Name(), "的权限是:", v.Mode())
		fmt.Println(v.Name(), "文件大小:", v.Size())
		fmt.Println(v.Name(), "创建时间", v.ModTime())
		fmt.Println(v.Name(), "系统信息", v.Sys())
		if v.IsDir() == true {
			fmt.Println(v.Name(), "是目录")
		}
	}
}

func TestReadFile(t *testing.T){
	data, err := ioutil.ReadFile("D:/test/widua.go")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func TestReadAll(t *testing.T){
	reader := strings.NewReader("hello word widuu") //返回*strings.Reader
	fmt.Println(reflect.TypeOf(reader))
	data, _ := ioutil.ReadAll(reader)
	fmt.Println(string(data))
}

func TestNopCloser(t *testing.T){
	reader := strings.NewReader("hello word widuu") //返回*strings.Reader
	r := ioutil.NopCloser(reader)
	defer r.Close()
	fmt.Println(reflect.TypeOf(reader))
	data, _ := ioutil.ReadAll(reader)
	fmt.Println(string(data))
}

func TestTempDir(t *testing.T){
	dir, err := ioutil.TempDir("D:/test", "tmp")
	if err != nil {
		fmt.Println("常见临时目录失败")
		return
	}
	fmt.Println(dir)  //返回的是D:\test\tmp846626247 就是前边的prefix+随机数
}

func TestTempFile(t *testing.T){
	file, error := ioutil.TempFile("D:/test", "tmp")
	defer file.Close()
	if error != nil {
		fmt.Println("创建文件失败")
		return
	}
	file.WriteString("Hello word") //利用file指针的WriteString()详情见os.WriteString()
	filedata, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(filedata))
}

