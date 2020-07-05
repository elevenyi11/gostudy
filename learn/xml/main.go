package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Name    string `xml:"name"`
	Address string `xml:"address,omitempty"`
	Hobby   string `xml:"-"`
	Father  string `xml:"parent>father"`
	Monther string `xml:"parent>monther"`
	Note    string `xml:"note,attr"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	stu1 := Student{
		Name:  "hah",
		Hobby: "basketball",
	}

	newData, err := xml.MarshalIndent(stu1, " ", "   ")
	checkErr(err)
	fmt.Println(string(newData))

	err = ioutil.WriteFile("stu.xml", newData, 0644)
	checkErr(err)

	content, err := ioutil.ReadFile("stu.xml")
	stu2 := &Student{}

	err = xml.Unmarshal(content, stu2)
	checkErr(err)
	fmt.Println("stu2: %#v\n", stu2)
}
