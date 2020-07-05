package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}

	return g
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func main() {
	var appPath string
	flag.StringVar(&appPath, "app-path", "app-path", "")
	flag.Parse()
	fmt.Printf("App path: %s", appPath)

	for _, v := range os.Environ() { //获取全部系统环境变量 获取的是 key=val 的[]string
		str := strings.Split(v, "=")
		fmt.Printf("key=%s,val=%s \n", str[0], str[1])
	}

	goPath := os.Getenv("GOPATH")

	currentPath := goPath + "\\src\\GoStudy\\branch\\"
	fmt.Printf("env: %s", goPath)

	path, _ := getCurrentPath()
	fmt.Println(path)
	filename := currentPath + "/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(80),
		grade(99),
		grade(100),
	)
}
