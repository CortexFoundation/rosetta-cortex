package main

import (
	"fmt"
	"github.com/CortexFoundation/rosetta-cortex/proxy"
	"github.com/CortexFoundation/rosetta-cortex/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func main() {
	settings := server.Settings{}
	settings.Client = proxy.New()
	settings.Listen = "127.0.0.1:8080"
	settings.Network = &types.NetworkIdentifier{"Cortex", "Mainnet", nil}
	s, err := server.NewServer(settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.Start()
	if err != nil {
		fmt.Println(err)
	}
}
