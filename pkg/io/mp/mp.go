package mp

import (
	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/io"
	"github.com/ont-bizsuite/ddxf-sdk/market_place_contract"
	"github.com/ont-bizsuite/ddxf-sdk/split_policy_contract"
)

// PublishItemInput ...
type PublishItemInput struct {
	// mp contract address
	MPContract string
	// item id, should be new
	ItemID string
	// seller account
	Seller string
	// meta hash of item meta
	ItemMetaHash string
	// endpoint for item meta hash
	MetaEndpoint string
	// token template ids is generated by /ddxf/dtoken/create_template
	TokenTemplateIDs []string
	// token contract for each token
	TokenContracts []string
	// accountant contract for marketplace
	AccountantContract string
	// split contract for multiple owners
	SplitContract string
	// split polocy for multiple owners
	SplitPolicy split_policy_contract.SplitPolicyRegisterParam
	// fee of item
	Fee market_place_contract.Fee
	// stock of item
	Stock uint64
	// valid until this timestamp, in seconds
	ExpiredDate uint64
}

// PublishItemOutput ...
// swagger:model PublishItemOutput
type PublishItemOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// UpdateItemInput ...
type UpdateItemInput struct {
	// mp contract address
	MPContract string
	// item id
	ItemID string
	// seller account
	Seller string
	// meta hash of item meta
	ItemMetaHash string
	// endpoint for item meta hash
	MetaEndpoint string
	// token template ids is generated by /ddxf/dtoken/create_template
	TokenTemplateIDs []string
	// token contract for each token
	TokenContracts []string
	// accountant contract for marketplace
	AccountantContract string
	// split contract for multiple owners
	SplitContract string
	// split polocy for multiple owners
	SplitPolicy split_policy_contract.SplitPolicyRegisterParam
	// fee of item
	Fee market_place_contract.Fee
	// stock of item
	Stock uint64
	// valid until this timestamp, in seconds
	ExpiredDate uint64
}

// UpdateItemOutput ...
// swagger:model UpdateItemOutput
type UpdateItemOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// DeleteItemInput ...
type DeleteItemInput struct {
	// mp contract address
	MPContract string
	// item id
	ItemID string
}

// DeleteItemOutput ...
// swagger:model DeleteItemOutput
type DeleteItemOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// BuyItemInput ...
type BuyItemInput struct {
	// mp contract address
	MPContract string
	// buyer account
	Buyer string
	// payer account
	Payer string
	// item id
	ItemID string
	// # of items to buy
	N int
}

// BuyItemOutput ...
// swagger:model BuyItemOutput
type BuyItemOutput struct {
	io.BaseResp
	// send with ontology sdk
	Tx string
}

// GetItemInput ...
type GetItemInput struct {
	// mp contract address
	MPContract string
	// item id
	ItemID string
}

// GetItemOutput ...
// swagger:model GetItemOutput
type GetItemOutput struct {
	io.BaseResp
	// meta id of item meta
	ItemMetaHash string
	// endpoint for item meta hash
	MetaEndpoint string
	// token template ids is generated by /ddxf/dtoken/create_template
	TokenTemplateIDs []string
	// token contract for each token
	TokenContracts []string
	// accountant contract for marketplace
	AccountantContract string
	// split polocy for multiple owners
	SplitPolicy split_policy_contract.SplitPolicyRegisterParam
	// fee of item
	Fee market_place_contract.Fee
	// stock of item
	Stock uint32
	// valid until this timestamp, in seconds
	ExpiredDate uint64
}
