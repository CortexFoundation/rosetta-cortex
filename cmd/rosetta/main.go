package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CortexFoundation/rosetta-cortex/router"
	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	serverPort = 8080
	url        = "http://127.0.0.1:8545"
)

func main() {
	network := &types.NetworkIdentifier{
		Blockchain: "Cortex",
		Network:    "Mainnet",
	}

	// The asserter automatically rejects incorrectly formatted
	// requests.
	asserter, err := asserter.NewServer(
		[]string{"Transfer", "Reward"},
		false,
		[]*types.NetworkIdentifier{network},
		nil,
		false,
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create the main router handler then apply the logger and Cors
	// middlewares in sequence.
	router := router.NewBlockchainRouter(network, asserter, url)

	loggedRouter := server.LoggerMiddleware(router)
	corsRouter := server.CorsMiddleware(loggedRouter)

	log.Printf("Listening on port %d\n", serverPort)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverPort),
		Handler:      corsRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
