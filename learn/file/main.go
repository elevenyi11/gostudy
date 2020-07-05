package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	bufferWriteFile()
}

func batchUpdateFileName() {
	dir := `D:\English\功夫英语\2.听读速成发音训练 Facefonics  ok\facefonics 修复`
	//dir_list, e := ioutil.ReadDir(dir)
	//if e != nil {
	//fmt.Println("read dir error")
	//return
	//}
	//for i, v := range dir_list {
	//fmt.Println(i, "=", v.Name())
	//}

	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
			//fmt.Println(path)
			if strings.Contains(info.Name(), "英文") {

				newFilePath := strings.Replace(path, "集", "集3", 1)
				fmt.Println(newFilePath)
				//os.Rename(path, newFilePath)
			}
		}
		return nil
	})

	time.Sleep(time.Second * 3)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile() {
	path := "test.txt"
	newFile, err := os.Create(path)
	checkError(err)
	defer newFile.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file doesn't exist!")
			return
		}
	}
	fmt.Println("file does exist, file name ", fileInfo.Name())
}

func readFile(path string) {
	data, err := ioutil.ReadFile(path)
	checkError(err)
	fmt.Println("file Content ", string(data))
}

func iWriteFile() {
	path := "test.txt"
	str := "hello"
	newStr := "world"
	newFile, err := os.Create(path)
	checkError(err)

	n1, err := newFile.WriteString(str)
	fmt.Println("n1: ", n1)
	readFile(path)

	n2, err := newFile.WriteAt([]byte(newStr), 6)
	checkError(err)

	fmt.Println("n2: ", n2)
	readFile(path)

	n3, err := newFile.WriteAt([]byte(newStr), 0)
	checkError(err)
	fmt.Println("n3: ", n3)
	readFile(path)

}

func bufferWriteFile() {
	path := "test.txt"
	str := "hello"

	newFile, err := os.Create(path)
	checkError(err)
	defer newFile.Close()

	bufferWrite := bufio.NewWriter(newFile)

	for _, v := range str {
		written, err := bufferWrite.WriteString(string(v))
		checkError(err)
		fmt.Println("written: ", written)
	}

	readFile(path)

	unflushSize := bufferWrite.Buffered()
	fmt.Println("unflushSize: ", unflushSize)

	bufferWrite.Flush()
	readFile(path)
}
