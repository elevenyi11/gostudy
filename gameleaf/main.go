package main

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func main() {
	Processor.Register(&Hello{})
}

type Hello struct {
	Name string
}
