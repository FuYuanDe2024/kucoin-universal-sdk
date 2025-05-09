// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package transfer

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// SubAccountTransferResp struct for SubAccountTransferResp
type SubAccountTransferResp struct {
	// common response
	CommonResponse *types.RestResponse
	// Transfer order ID
	OrderId string `json:"orderId,omitempty"`
}

// NewSubAccountTransferResp instantiates a new SubAccountTransferResp object
// This constructor will assign default values to properties that have it defined
func NewSubAccountTransferResp(orderId string) *SubAccountTransferResp {
	this := SubAccountTransferResp{}
	this.OrderId = orderId
	return &this
}

// NewSubAccountTransferRespWithDefaults instantiates a new SubAccountTransferResp object
// This constructor will only assign default values to properties that have it defined,
func NewSubAccountTransferRespWithDefaults() *SubAccountTransferResp {
	this := SubAccountTransferResp{}
	return &this
}

func (o *SubAccountTransferResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize
}

func (o *SubAccountTransferResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
