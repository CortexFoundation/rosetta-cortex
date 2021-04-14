package services

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// NetworkAPIService implements the server.NetworkAPIServicer interface.
type NetworkAPIService struct {
	network *types.NetworkIdentifier
}

// NewNetworkAPIService creates a new instance of a NetworkAPIService.
func NewNetworkAPIService(network *types.NetworkIdentifier) server.NetworkAPIServicer {
	return &NetworkAPIService{
		network: network,
	}
}

// NetworkList implements the /network/list endpoint
func (s *NetworkAPIService) NetworkList(
	ctx context.Context,
	request *types.MetadataRequest,
) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			s.network,
		},
	}, nil
}

// NetworkStatus implements the /network/status endpoint.
func (s *NetworkAPIService) NetworkStatus(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkStatusResponse, *types.Error) {
	//TODO
	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "0xec87df31c230298a66eabbfa3d030a835831a55ddbefdc958e77e2f7cd59e81d",
		},
		CurrentBlockTimestamp: int64(1618217080000),
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Index: 0,
			Hash:  "0x21d6ce908e2d1464bd74bbdbf7249845493cc1ba10460758169b978e187762c1",
		},
		Peers: []*types.Peer{
			{
				PeerID: "0d62a13d5f696e1b0b97597ef475494d3aae0eeca85a8cd22036d403f7b3ef90",
			},
		},
	}, nil
}

// NetworkOptions implements the /network/options endpoint.
func (s *NetworkAPIService) NetworkOptions(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkOptionsResponse, *types.Error) {
	//TODO
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.4.0",
			NodeVersion:    "v1.10.22",
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     "Success",
					Successful: true,
				},
				{
					Status:     "Reverted",
					Successful: false,
				},
			},
			OperationTypes: []string{
				"Transfer",
				"Reward",
			},
			Errors: []*types.Error{
				{
					Code:      1,
					Message:   "not implemented",
					Retriable: false,
				},
			},
		},
	}, nil
}
