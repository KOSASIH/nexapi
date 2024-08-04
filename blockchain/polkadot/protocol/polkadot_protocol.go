package protocol

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// PolkadotProtocol represents the Polkadot protocol
type PolkadotProtocol struct {
	nodeURL string
	client  *types.Client
}

// NewPolkadotProtocol creates a new PolkadotProtocol instance
func NewPolkadotProtocol(nodeURL string) (*PolkadotProtocol, error) {
	client, err := types.NewClient(nodeURL)
	if err!= nil {
		return nil, err
	}
	return &PolkadotProtocol{nodeURL: nodeURL, client: client}, nil
}

// GetBlockHash returns the block hash at a given height
func (p *PolkadotProtocol) GetBlockHash(height uint64) (string, error) {
	hash, err := p.client.RPC.BlockHash(context.Background(), height)
	if err!= nil {
		return "", err
	}
	return hash.Hex(), nil
}

// GetBlock returns the block at a given hash
func (p *PolkadotProtocol) GetBlock(hash string) (*types.Block, error) {
	block, err := p.client.RPC.Block(context.Background(), hash)
	if err!= nil {
		return nil, err
	}
	return block, nil
}

// GetFinalizedHead returns the finalized head of the chain
func (p *PolkadotProtocol) GetFinalizedHead() (*types.Header, error) {
	header, err := p.client.RPC.FinalizedHead(context.Background())
	if err!= nil {
		return nil, err
	}
	return header, nil
}

// GetBestNumber returns the best block number
func (p *PolkadotProtocol) GetBestNumber() (uint64, error) {
	number, err := p.client.RPC.BestNumber(context.Background())
	if err!= nil {
		return 0, err
	}
	return number, nil
}

// GetMetadata returns the metadata of the chain
func (p *PolkadotProtocol) GetMetadata() (*types.Metadata, error) {
	metadata, err := p.client.RPC.Metadata(context.Background())
	if err!= nil {
		return nil, err
	}
	return metadata, nil
}

// GetStorage returns the storage value at a given key
func (p *PolkadotProtocol) GetStorage(key string) (string, error) {
	data, err := p.client.RPC.Storage(context.Background(), key)
	if err!= nil {
		return "", err
	}
	return data.Hex(), nil
}

// ExtrinsicSign signs an extrinsic with a given signer
func (p *PolkadotProtocol) ExtrinsicSign(ext *types.Extrinsic, signer *ecdsa.PrivateKey) (*types.Extrinsic, error) {
	signedExt, err := ext.Sign(signer)
	if err!= nil {
		return nil, err
	}
	return signedExt, nil
}

// SubmitExtrinsic submits an extrinsic to the network
func (p *PolkadotProtocol) SubmitExtrinsic(ext *types.Extrinsic) (string, error) {
	hash, err := p.client.RPC.SubmitExtrinsic(context.Background(), ext)
	if err!= nil {
		return "", err
	}
	return hash.Hex(), nil
}

// QuerySubscribe subscribes to a query
func (p *PolkadotProtocol) QuerySubscribe(query *types.Query) (*types.Subscriber, error) {
	subscriber, err := p.client.RPC.QuerySubscribe(context.Background(), query)
	if err!= nil {
		return nil, err
	}
	return subscriber, nil
}
