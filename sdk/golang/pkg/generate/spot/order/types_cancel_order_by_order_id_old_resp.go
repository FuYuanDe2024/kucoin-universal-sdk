// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// CancelOrderByOrderIdOldResp struct for CancelOrderByOrderIdOldResp
type CancelOrderByOrderIdOldResp struct {
	// common response
	CommonResponse    *types.RestResponse
	CancelledOrderIds []string `json:"cancelledOrderIds,omitempty"`
}

// NewCancelOrderByOrderIdOldResp instantiates a new CancelOrderByOrderIdOldResp object
// This constructor will assign default values to properties that have it defined
func NewCancelOrderByOrderIdOldResp(cancelledOrderIds []string) *CancelOrderByOrderIdOldResp {
	this := CancelOrderByOrderIdOldResp{}
	this.CancelledOrderIds = cancelledOrderIds
	return &this
}

// NewCancelOrderByOrderIdOldRespWithDefaults instantiates a new CancelOrderByOrderIdOldResp object
// This constructor will only assign default values to properties that have it defined,
func NewCancelOrderByOrderIdOldRespWithDefaults() *CancelOrderByOrderIdOldResp {
	this := CancelOrderByOrderIdOldResp{}
	return &this
}

func (o *CancelOrderByOrderIdOldResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["cancelledOrderIds"] = o.CancelledOrderIds
	return toSerialize
}

func (o *CancelOrderByOrderIdOldResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
