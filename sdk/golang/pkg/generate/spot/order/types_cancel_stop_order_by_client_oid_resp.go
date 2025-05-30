// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// CancelStopOrderByClientOidResp struct for CancelStopOrderByClientOidResp
type CancelStopOrderByClientOidResp struct {
	// common response
	CommonResponse *types.RestResponse
	// Client Order Id，unique identifier created by the user
	ClientOid string `json:"clientOid,omitempty"`
	// Unique ID of the cancelled order
	CancelledOrderId string `json:"cancelledOrderId,omitempty"`
}

// NewCancelStopOrderByClientOidResp instantiates a new CancelStopOrderByClientOidResp object
// This constructor will assign default values to properties that have it defined
func NewCancelStopOrderByClientOidResp(clientOid string, cancelledOrderId string) *CancelStopOrderByClientOidResp {
	this := CancelStopOrderByClientOidResp{}
	this.ClientOid = clientOid
	this.CancelledOrderId = cancelledOrderId
	return &this
}

// NewCancelStopOrderByClientOidRespWithDefaults instantiates a new CancelStopOrderByClientOidResp object
// This constructor will only assign default values to properties that have it defined,
func NewCancelStopOrderByClientOidRespWithDefaults() *CancelStopOrderByClientOidResp {
	this := CancelStopOrderByClientOidResp{}
	return &this
}

func (o *CancelStopOrderByClientOidResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["clientOid"] = o.ClientOid
	toSerialize["cancelledOrderId"] = o.CancelledOrderId
	return toSerialize
}

func (o *CancelStopOrderByClientOidResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
