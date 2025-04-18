// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package subaccount

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// AddSubAccountFuturesPermissionResp struct for AddSubAccountFuturesPermissionResp
type AddSubAccountFuturesPermissionResp struct {
	// common response
	CommonResponse *types.RestResponse
	Data           bool `json:"data,omitempty"`
}

// NewAddSubAccountFuturesPermissionResp instantiates a new AddSubAccountFuturesPermissionResp object
// This constructor will assign default values to properties that have it defined
func NewAddSubAccountFuturesPermissionResp(data bool) *AddSubAccountFuturesPermissionResp {
	this := AddSubAccountFuturesPermissionResp{}
	this.Data = data
	return &this
}

// NewAddSubAccountFuturesPermissionRespWithDefaults instantiates a new AddSubAccountFuturesPermissionResp object
// This constructor will only assign default values to properties that have it defined,
func NewAddSubAccountFuturesPermissionRespWithDefaults() *AddSubAccountFuturesPermissionResp {
	this := AddSubAccountFuturesPermissionResp{}
	return &this
}

func (o *AddSubAccountFuturesPermissionResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize
}

func (o *AddSubAccountFuturesPermissionResp) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.Data)
	return err
}

func (o *AddSubAccountFuturesPermissionResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
