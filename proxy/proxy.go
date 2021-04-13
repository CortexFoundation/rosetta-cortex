package proxy

import (
	"context"
	//mocks "github.com/CortexFoundation/rosetta-cortex/ethereum"
	ctypes "github.com/CortexFoundation/rosetta-cortex/types"
	"github.com/coinbase/rosetta-sdk-go/types"
)

type Proxy struct {
	c *mocks.Client
}

func New() *Proxy {
	proxy := &Proxy{}
	proxy.c = mocks.NewClient("127.0.0.1:8545", nil, nil)
	return proxy
}

// Needed if the client needs to perform some action before connecting
func (oc *Proxy) Bootstrap() error {
	return nil
}

// Ready checks if the servicer constraints for queries are satisfied
// for example the node might still not be ready, it's useful in process
// when the rosetta instance might come up before the node itself
// the servicer must return nil if the node is ready
func (oc *Proxy) Ready() error {
	return nil
}

// Data API

// Balances fetches the balance of the given address
// if height is not nil, then the balance will be displayed
// at the provided height, otherwise last block balance will be returned
func (oc *Proxy) Balances(ctx context.Context, addr string, height *int64) ([]*types.Amount, error) {
	return nil, nil
}

// BlockByHashAlt gets a block and its transaction at the provided height
func (oc *Proxy) BlockByHash(ctx context.Context, hash string) (ctypes.BlockResponse, error) {
	return ctypes.BlockResponse{}, nil
}

// BlockByHeightAlt gets a block given its height, if height is nil then last block is returned
func (oc *Proxy) BlockByHeight(ctx context.Context, height *int64) (ctypes.BlockResponse, error) {
	return ctypes.BlockResponse{}, nil
}

// BlockTransactionsByHash gets the block, parent block and transactions
// given the block hash.
func (oc *Proxy) BlockTransactionsByHash(ctx context.Context, hash string) (ctypes.BlockTransactionsResponse, error) {
	return ctypes.BlockTransactionsResponse{}, nil
}

// BlockTransactionsByHash gets the block, parent block and transactions
// given the block hash.
func (oc *Proxy) BlockTransactionsByHeight(ctx context.Context, height *int64) (ctypes.BlockTransactionsResponse, error) {
	return ctypes.BlockTransactionsResponse{}, nil
}

// GetTx gets a transaction given its hash
func (oc *Proxy) GetTx(ctx context.Context, hash string) (*types.Transaction, error) {
	return nil, nil
}

// GetUnconfirmedTx gets an unconfirmed Tx given its hash
// NOTE(fdymylja): NOT IMPLEMENTED YET!
func (oc *Proxy) GetUnconfirmedTx(ctx context.Context, hash string) (*types.Transaction, error) {
	return nil, nil
}

// Mempool returns the list of the current non confirmed transactions
func (oc *Proxy) Mempool(ctx context.Context) ([]*types.TransactionIdentifier, error) {
	return nil, nil
}

// Peers gets the peers currently connected to the node
func (oc *Proxy) Peers(ctx context.Context) ([]*types.Peer, error) {
	return nil, nil
}

// Status returns the node status, such as sync data, version etc
func (oc *Proxy) Status(ctx context.Context) (*types.SyncStatus, error) {
	return nil, nil
}

// Construction API

// PostTx posts txBytes to the node and returns the transaction identifier plus metadata related
// to the transaction itself.
func (oc *Proxy) PostTx(txBytes []byte) (res *types.TransactionIdentifier, meta map[string]interface{}, err error) {
	return nil, nil, nil
}

// ConstructionMetadataFromOptions
func (oc *Proxy) ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error) {
	return nil, nil
}

// SupportedOperations lists the operations supported by the implementation
func (oc *Proxy) SupportedOperations() []string {
	return []string{"Block"}
}

// OperationsStatuses returns the list of statuses supported by the implementation
func (oc *Proxy) OperationStatuses() []*types.OperationStatus {
	return nil
}

// Version returns the version of the node
func (oc *Proxy) Version() string {
	return ""
}

// SignedTx returns the signed transaction given the tx bytes (msgs) plus the signatures
func (oc *Proxy) SignedTx(ctx context.Context, txBytes []byte, sigs []*types.Signature) (signedTxBytes []byte, err error) {
	return nil, nil
}

// TxOperationsAndSignersAccountIdentifiers returns the operations related to a transaction and the account
// identifiers if the transaction is signed
func (oc *Proxy) TxOperationsAndSignersAccountIdentifiers(signed bool, hexBytes []byte) (ops []*types.Operation, signers []*types.AccountIdentifier, err error) {
	return nil, nil, nil
}

// ConstructionPayload returns the construction payload given the request
func (oc *Proxy) ConstructionPayload(ctx context.Context, req *types.ConstructionPayloadsRequest) (resp *types.ConstructionPayloadsResponse, err error) {
	return nil, nil
}

// PreprocessOperationsToOptions returns the options given the preprocess operations
func (oc *Proxy) PreprocessOperationsToOptions(ctx context.Context, req *types.ConstructionPreprocessRequest) (resp *types.ConstructionPreprocessResponse, err error) {
	return nil, nil
}

// AccountIdentifierFromPublicKey returns the account identifier given the public key
func (oc *Proxy) AccountIdentifierFromPublicKey(pubKey *types.PublicKey) (*types.AccountIdentifier, error) {
	return nil, nil
}
