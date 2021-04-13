package services

import (
	"context"
	"fmt"
	"time"

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

// Block implements the /block endpoint.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "block 1000",
		},
		Balances: []*Amount{&types.Amount{Value: "1000", Currency: &types.Currency{Symbol: "CTXC", Decimals: 2}}},
		Metadata: nil,
	}
}

// BlockTransaction implements the /block/transaction endpoint.
func (s *BlockAPIService) BlockTransaction(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {
	return &types.AccountCoinsResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "block 1000",
		},
		Coins:    nil,
		Metadata: nil,
	}
}
