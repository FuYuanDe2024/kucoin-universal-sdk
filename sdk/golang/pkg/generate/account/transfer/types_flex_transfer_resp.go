// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package transfer

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// FlexTransferResp struct for FlexTransferResp
type FlexTransferResp struct {
	// common response
	CommonResponse *types.RestResponse
	// Transfer order ID
	OrderId string `json:"orderId,omitempty"`
}

// NewFlexTransferResp instantiates a new FlexTransferResp object
// This constructor will assign default values to properties that have it defined
func NewFlexTransferResp(orderId string) *FlexTransferResp {
	this := FlexTransferResp{}
	this.OrderId = orderId
	return &this
}

// NewFlexTransferRespWithDefaults instantiates a new FlexTransferResp object
// This constructor will only assign default values to properties that have it defined,
func NewFlexTransferRespWithDefaults() *FlexTransferResp {
	this := FlexTransferResp{}
	return &this
}

func (o *FlexTransferResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize
}

func (o *FlexTransferResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
