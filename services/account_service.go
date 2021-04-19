package services

import (
	"context"

	"github.com/CortexFoundation/rosetta-cortex/proxy"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// BlockAPIService implements the server.BlockAPIServicer interface.
type AccountAPIService struct {
	network *types.NetworkIdentifier
	proxy   *proxy.Proxy
}

// NewBlockAPIService creates a new instance of a BlockAPIService.
func NewAccountAPIService(network *types.NetworkIdentifier, p *proxy.Proxy) server.AccountAPIServicer {
	return &AccountAPIService{
		network: network,
		proxy:   p,
	}
}

// Block implements the /account/balance endpoint.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "0xec87df31c230298a66eabbfa3d030a835831a55ddbefdc958e77e2f7cd59e81d",
		},
		Balances: []*types.Amount{
			&types.Amount{
				Value:    "1000",
				Currency: &types.Currency{Symbol: "CTXC", Decimals: 18},
			},
		},
	}, nil
}

// BlockTransaction implements the /account/coins endpoint
func (s *AccountAPIService) AccountCoins(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {
	return &types.AccountCoinsResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "0xec87df31c230298a66eabbfa3d030a835831a55ddbefdc958e77e2f7cd59e81d",
		},
		Coins: []*types.Coin{
			{
				Amount: &types.Amount{
					Value: "1000",
					Currency: &types.Currency{
						Symbol:   "CTXC",
						Decimals: 18,
					},
				},
				CoinIdentifier: &types.CoinIdentifier{
					Identifier: "coin",
				},
			},
		},
	}, nil
}
