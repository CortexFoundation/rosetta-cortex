package router

import (
	"net/http"

	"github.com/CortexFoundation/rosetta-cortex/proxy"
	"github.com/CortexFoundation/rosetta-cortex/services"
	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
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
