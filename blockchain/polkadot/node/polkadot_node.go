package node

import (
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/interfaces"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/primitive"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/events"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/query"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/rpc"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/sr25519"
	"github.com/centrifuge/go-substrate-rpc-client/v4/websocket"
)

// PolkadotNode represents a Polkadot node
type PolkadotNode struct {
	*websocket.Client
}

// NewPolkadotNode creates a new Polkadot node
func NewPolkadotNode(url string) (*PolkadotNode, error) {
	client, err := websocket.New(url, nil)
	if err != nil {
		return nil, err
	}
	return &PolkadotNode{client}, nil
}

// GetBlockHash gets the hash of a block by its height
func (n *PolkadotNode) GetBlockHash(height uint64) (types.Hash, error) {
	return n.Client.Chain.GetBlockHash(height)
}

// GetBlock gets a block by its hash
func (n *PolkadotNode) GetBlock(hash types.Hash) (types.Block, error) {
	return n.Client.Chain.GetBlock(hash)
}

// GetFinalizedHead gets the finalized head of the chain
func (n *PolkadotNode) GetFinalizedHead() (types.Header, error) {
	return n.Client.Chain.GetFinalizedHead()
}

// GetBestNumber gets the best block number
func (n *PolkadotNode) GetBestNumber() (uint64, error) {
	return n.Client.Chain.GetBestNumber()
}

// GetMetadata gets the metadata of the chain
func (n *PolkadotNode) GetMetadata() (types.Metadata, error) {
	return n.Client.State.GetMetadata()
}

// GetStorage gets the value of a storage item
func (n *PolkadotNode) GetStorage(item types.StorageKey) (types.StorageData, error) {
	return n.Client.State.GetStorage(item)
}

// GetEventRecords gets the event records for a block
func (n *PolkadotNode) GetEventRecords(hash types.Hash, event types.Event) ([]events.EventRecord, error) {
	return n.Client.Events.GetEventRecords(hash, event)
}

// ExtrinsicSign signs an extrinsic
func (n *PolkadotNode) ExtrinsicSign(ext extrinsic.Extrinsic, signer sr25519.Pair) (extrinsic.Extrinsic, error) {
	return ext.Sign(signer)
}

// SubmitExtrinsic submits an extrinsic
func (n *PolkadotNode) SubmitExtrinsic(ext extrinsic.Extrinsic) (types.Hash, error) {
	return n.Client.Extrinsic.Submit(ext)
}

// QuerySubscribe subscribes to a query
func (n *PolkadotNode) QuerySubscribe(query query.Subscription) (query.Subscriber, error) {
	return n.Client.Subscribe(query)
}
