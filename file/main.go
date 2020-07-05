package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var fileFolderPath = `F:\迅雷下载\BreachCompilation\data`

func main() {
	files, dirs, _ := GetFilesAndDirs(fileFolderPath)

	for _, dir := range dirs {
		fmt.Printf("获取的文件夹为[%s]\n", dir)
	}
	fmt.Printf("=======================================\n")
	var totalSize int64
	for _, table := range dirs {
		temp, size, _ := GetAllFiles(table)
		totalSize += size
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	for _, table1 := range files {
		fmt.Printf("获取的文件为[%s]\n", table1)
	}
	fmt.Printf("=======================================\n")
	fmt.Printf("获取的文件总数为[%s], size:[%d]\n", len(files), totalSize/1024/1024/1024)
}

func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

func GetAllFiles(dirPth string) (files []string, totalSize int64, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, 0, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			_, size, _ := GetAllFiles(dirPth + PthSep + fi.Name())
			totalSize += size
		} else {
			fiSize := strconv.FormatInt(fi.Size(), 10)
			totalSize += fi.Size()
			files = append(files, dirPth+PthSep+fi.Name()+" size: "+fiSize)
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, size, _ := GetAllFiles(table)
		totalSize += size
		for _, temp1 := range temp {
			fi, err := os.Stat(temp1)
			if err == nil {
				totalSize += fi.Size()
				fiSize := strconv.FormatInt(fi.Size(), 10)
				files = append(files, temp1+" size: "+fiSize)
			}

		}
	}

	return files, totalSize, nil
}
