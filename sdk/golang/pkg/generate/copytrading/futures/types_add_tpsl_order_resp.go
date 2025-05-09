// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package futures

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// AddTPSLOrderResp struct for AddTPSLOrderResp
type AddTPSLOrderResp struct {
	// common response
	CommonResponse *types.RestResponse
	// The unique order id generated by the trading system,which can be used later for further actions such as canceling the order.
	OrderId string `json:"orderId,omitempty"`
	// The user self-defined order id.
	ClientOid string `json:"clientOid,omitempty"`
}

// NewAddTPSLOrderResp instantiates a new AddTPSLOrderResp object
// This constructor will assign default values to properties that have it defined
func NewAddTPSLOrderResp(orderId string, clientOid string) *AddTPSLOrderResp {
	this := AddTPSLOrderResp{}
	this.OrderId = orderId
	this.ClientOid = clientOid
	return &this
}

// NewAddTPSLOrderRespWithDefaults instantiates a new AddTPSLOrderResp object
// This constructor will only assign default values to properties that have it defined,
func NewAddTPSLOrderRespWithDefaults() *AddTPSLOrderResp {
	this := AddTPSLOrderResp{}
	return &this
}

func (o *AddTPSLOrderResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	toSerialize["clientOid"] = o.ClientOid
	return toSerialize
}

func (o *AddTPSLOrderResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
