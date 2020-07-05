package main

import (
	"reflect"

	"github.com/name5566/leaf/module"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	module.Skeleton{}.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*Hello)

	a := args[1].(gate.Agent)

	log.Debug("hello v%", m.Name)

	a.WriteMsg(&Hello{Name: "client"})
}
