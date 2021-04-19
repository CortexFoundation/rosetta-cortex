package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CortexFoundation/rosetta-cortex/proxy"
	"github.com/CortexFoundation/rosetta-cortex/services"
	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	serverPort = 8080
)

// NewBlockchainRouter creates a Mux http.Handler from a collection
// of server controllers.
func NewBlockchainRouter(
	network *types.NetworkIdentifier,
	asserter *asserter.Asserter,
) http.Handler {

	url := "http://127.0.0.1:8545"
	p := proxy.New(url)

	networkAPIService := services.NewNetworkAPIService(network, p)
	networkAPIController := server.NewNetworkAPIController(
		networkAPIService,
		asserter,
	)

	blockAPIService := services.NewBlockAPIService(network, p)
	blockAPIController := server.NewBlockAPIController(
		blockAPIService,
		asserter,
	)

	accountAPIService := services.NewAccountAPIService(network, p)
	accountAPIController := server.NewAccountAPIController(
		accountAPIService,
		asserter,
	)

	return server.NewRouter(networkAPIController, blockAPIController, accountAPIController)
}

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
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create the main router handler then apply the logger and Cors
	// middlewares in sequence.
	router := NewBlockchainRouter(network, asserter)
	loggedRouter := server.LoggerMiddleware(router)
	corsRouter := server.CorsMiddleware(loggedRouter)
	log.Printf("Listening on port %d\n", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), corsRouter))
}
