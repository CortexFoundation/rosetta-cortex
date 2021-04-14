package services

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// BlockAPIService implements the server.BlockAPIServicer interface.
type AccountAPIService struct {
	network *types.NetworkIdentifier
}

// NewBlockAPIService creates a new instance of a BlockAPIService.
func NewAccountAPIService(network *types.NetworkIdentifier) server.AccountAPIServicer {
	return &AccountAPIService{
		network: network,
	}
}

// Block implements the /account/balance endpoint.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 3616244,
			Hash:  "0xec87df31c230298a66eabbfa3d030a835831a55ddbefdc958e77e2f7cd59e81d",
		},
		Balances: []*types.Amount{&types.Amount{Value: "1000", Currency: &types.Currency{Symbol: "CTXC", Decimals: 2}}},
		Metadata: nil,
	}, nil
}

// BlockTransaction implements the /account/coins endpointfunc (s *AccountAPIService) AccountBalance(
func (s *AccountAPIService) AccountCoins(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {
	return &types.AccountCoinsResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 3616244,
			Hash:  "0xec87df31c230298a66eabbfa3d030a835831a55ddbefdc958e77e2f7cd59e81d",
		},
		Coins:    nil,
		Metadata: nil,
	}, nil
}
