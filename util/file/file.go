/*
Copyright 2017 by Eleven.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//GetBinaryCurrentPath get current binary path
func GetBinaryCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	if strings.Contains(path, "command-line-arguments") {
		return GetCurrentPath()
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

// GetCurrentPath get current execute command path
func GetCurrentPath() (string, error) {
	return os.Getwd()
}

// MakeDir Create dir by the file name
func MakeDir(fileDir string) error {
	return os.MkdirAll(fileDir, 0777)
}

// MakeDirByFile Create dir by the filename
// ./dir/filename  /home/dir/filename
func MakeDirByFile(filePath string) error {
	temp := strings.Split(filePath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirPath := strings.Join(temp[0:len(temp)+1], "/")
	return MakeDir(dirPath)
}

//FileExists check file is exist
// if file exist return true, otherwise return false
func FileExists(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}

	if fi.IsDir() {
		return false
	}

	return true
}

// ReadFromFile base ioutil ReadFile function
func ReadFromFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// GetFilenameInfo base os.Stat
func GetFilenameInfo(filepath string) (os.FileInfo, error) {
	fileinfo, err := os.Stat(filepath)
	return fileinfo, err
}

//SaveToFile
func SaveToFile(filepath string, content []byte) error {
	//all right write to file
	err := ioutil.WriteFile(filepath, content, 0777)
	return err
}

// rename
func Rename(oldfilename string, newfilename string) error {
	return os.Rename(oldfilename, newfilename)
}

//OpenFile
func OpenFile(filePath string) (*os.File, error) {
	outFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Can not open file %s", filePath)
	}
	return outFile, err
}
