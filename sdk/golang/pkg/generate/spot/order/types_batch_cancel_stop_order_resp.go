// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// BatchCancelStopOrderResp struct for BatchCancelStopOrderResp
type BatchCancelStopOrderResp struct {
	// common response
	CommonResponse *types.RestResponse
	// order id array
	CancelledOrderIds []string `json:"cancelledOrderIds,omitempty"`
}

// NewBatchCancelStopOrderResp instantiates a new BatchCancelStopOrderResp object
// This constructor will assign default values to properties that have it defined
func NewBatchCancelStopOrderResp(cancelledOrderIds []string) *BatchCancelStopOrderResp {
	this := BatchCancelStopOrderResp{}
	this.CancelledOrderIds = cancelledOrderIds
	return &this
}

// NewBatchCancelStopOrderRespWithDefaults instantiates a new BatchCancelStopOrderResp object
// This constructor will only assign default values to properties that have it defined,
func NewBatchCancelStopOrderRespWithDefaults() *BatchCancelStopOrderResp {
	this := BatchCancelStopOrderResp{}
	return &this
}

func (o *BatchCancelStopOrderResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["cancelledOrderIds"] = o.CancelledOrderIds
	return toSerialize
}

func (o *BatchCancelStopOrderResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
