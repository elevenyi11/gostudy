package main

import (
	"leafserver/src/server/game"
)

func init() {
	Processor.SetRouter(&Hello{}, game.ChanRPC)
}
