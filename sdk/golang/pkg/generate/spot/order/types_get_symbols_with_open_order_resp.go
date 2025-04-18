// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetSymbolsWithOpenOrderResp struct for GetSymbolsWithOpenOrderResp
type GetSymbolsWithOpenOrderResp struct {
	// common response
	CommonResponse *types.RestResponse
	// The symbol that has active orders
	Symbols []string `json:"symbols,omitempty"`
}

// NewGetSymbolsWithOpenOrderResp instantiates a new GetSymbolsWithOpenOrderResp object
// This constructor will assign default values to properties that have it defined
func NewGetSymbolsWithOpenOrderResp(symbols []string) *GetSymbolsWithOpenOrderResp {
	this := GetSymbolsWithOpenOrderResp{}
	this.Symbols = symbols
	return &this
}

// NewGetSymbolsWithOpenOrderRespWithDefaults instantiates a new GetSymbolsWithOpenOrderResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetSymbolsWithOpenOrderRespWithDefaults() *GetSymbolsWithOpenOrderResp {
	this := GetSymbolsWithOpenOrderResp{}
	return &this
}

func (o *GetSymbolsWithOpenOrderResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["symbols"] = o.Symbols
	return toSerialize
}

func (o *GetSymbolsWithOpenOrderResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
