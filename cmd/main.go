package main

import (
	"github.com/CortexFoundation/rosetta-cortex/proxy"
	"github.com/CortexFoundation/rosetta-cortex/server"
)

func main() {
	settings := server.Settings{}
	settings.Client = proxy.New()
	s, _ := server.NewServer(settings)
	s.Start()
}
