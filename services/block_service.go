package services

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// BlockAPIService implements the server.BlockAPIServicer interface.
type BlockAPIService struct {
	network *types.NetworkIdentifier
}

// NewBlockAPIService creates a new instance of a BlockAPIService.
func NewBlockAPIService(network *types.NetworkIdentifier) server.BlockAPIServicer {
	return &BlockAPIService{
		network: network,
	}
}

// Block implements the /block endpoint.
func (s *BlockAPIService) Block(
	ctx context.Context,
	request *types.BlockRequest,
) (*types.BlockResponse, *types.Error) {
	if *request.BlockIdentifier.Index != 1000 {
		previousBlockIndex := *request.BlockIdentifier.Index - 1
		if previousBlockIndex < 0 {
			previousBlockIndex = 0
		}

		return &types.BlockResponse{
			Block: &types.Block{
				BlockIdentifier: &types.BlockIdentifier{
					Index: *request.BlockIdentifier.Index,
					Hash:  fmt.Sprintf("block %d", *request.BlockIdentifier.Index),
				},
				ParentBlockIdentifier: &types.BlockIdentifier{
					Index: previousBlockIndex,
					Hash:  fmt.Sprintf("block %d", previousBlockIndex),
				},
				Timestamp:    time.Now().UnixNano() / 1000000,
				Transactions: []*types.Transaction{},
			},
		}, nil
	}

	//TODO

	return &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier: &types.BlockIdentifier{
				Index: 1000,
				Hash:  "block 1000",
			},
			ParentBlockIdentifier: &types.BlockIdentifier{
				Index: 999,
				Hash:  "block 999",
			},
			Timestamp: 1586483189000,
			Transactions: []*types.Transaction{
				{
					TransactionIdentifier: &types.TransactionIdentifier{
						Hash: "transaction 0",
					},
					Operations: []*types.Operation{
						{
							OperationIdentifier: &types.OperationIdentifier{
								Index: 0,
							},
							Type:   "Transfer",
							Status: types.String("Success"),
							Account: &types.AccountIdentifier{
								Address: "account 0",
							},
							Amount: &types.Amount{
								Value: "1000",
								Currency: &types.Currency{
									Symbol:   "CTXC",
									Decimals: 18,
								},
							},
						},
						{
							OperationIdentifier: &types.OperationIdentifier{
								Index: 1,
							},
							RelatedOperations: []*types.OperationIdentifier{
								{
									Index: 0,
								},
							},
							Type:   "Transfer",
							Status: types.String("Reverted"),
							Account: &types.AccountIdentifier{
								Address: "account 1",
							},
							Amount: &types.Amount{
								Value: "1000",
								Currency: &types.Currency{
									Symbol:   "CTXC",
									Decimals: 18,
								},
							},
						},
					},
				},
			},
		},
		OtherTransactions: []*types.TransactionIdentifier{
			{
				Hash: "transaction 1",
			},
		},
	}, nil
}

// BlockTransaction implements the /block/transaction endpoint.
func (s *BlockAPIService) BlockTransaction(
	ctx context.Context,
	request *types.BlockTransactionRequest,
) (*types.BlockTransactionResponse, *types.Error) {
	//TODO
	return &types.BlockTransactionResponse{
		Transaction: &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: "transaction 1",
			},
			Operations: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 0,
					},
					Type:   "Reward",
					Status: types.String("Success"),
					Account: &types.AccountIdentifier{
						Address: "account 1",
					},
					Amount: &types.Amount{
						Value: "1000",
						Currency: &types.Currency{
							Symbol:   "CTXC",
							Decimals: 18,
						},
					},
				},
			},
		},
	}, nil
}
