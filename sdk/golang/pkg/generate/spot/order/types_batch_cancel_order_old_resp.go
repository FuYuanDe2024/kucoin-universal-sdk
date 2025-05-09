// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// BatchCancelOrderOldResp struct for BatchCancelOrderOldResp
type BatchCancelOrderOldResp struct {
	// common response
	CommonResponse    *types.RestResponse
	CancelledOrderIds []string `json:"cancelledOrderIds,omitempty"`
}

// NewBatchCancelOrderOldResp instantiates a new BatchCancelOrderOldResp object
// This constructor will assign default values to properties that have it defined
func NewBatchCancelOrderOldResp(cancelledOrderIds []string) *BatchCancelOrderOldResp {
	this := BatchCancelOrderOldResp{}
	this.CancelledOrderIds = cancelledOrderIds
	return &this
}

// NewBatchCancelOrderOldRespWithDefaults instantiates a new BatchCancelOrderOldResp object
// This constructor will only assign default values to properties that have it defined,
func NewBatchCancelOrderOldRespWithDefaults() *BatchCancelOrderOldResp {
	this := BatchCancelOrderOldResp{}
	return &this
}

func (o *BatchCancelOrderOldResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["cancelledOrderIds"] = o.CancelledOrderIds
	return toSerialize
}

func (o *BatchCancelOrderOldResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
