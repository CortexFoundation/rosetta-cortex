package main

import (
	"github.com/CortexFoundation/rosetta-cortex/server"
)

func main() {
	settings := server.Settings{}
	s, _ := server.NewServer(settings)
	s.Start()
}
