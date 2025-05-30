// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package futures

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// ModifyIsolatedMarginRiskLimtResp struct for ModifyIsolatedMarginRiskLimtResp
type ModifyIsolatedMarginRiskLimtResp struct {
	// common response
	CommonResponse *types.RestResponse
	// To adjust the level will cancel the open order, the response can only indicate whether the submit of the adjustment request is successful or not.
	Data bool `json:"data,omitempty"`
}

// NewModifyIsolatedMarginRiskLimtResp instantiates a new ModifyIsolatedMarginRiskLimtResp object
// This constructor will assign default values to properties that have it defined
func NewModifyIsolatedMarginRiskLimtResp(data bool) *ModifyIsolatedMarginRiskLimtResp {
	this := ModifyIsolatedMarginRiskLimtResp{}
	this.Data = data
	return &this
}

// NewModifyIsolatedMarginRiskLimtRespWithDefaults instantiates a new ModifyIsolatedMarginRiskLimtResp object
// This constructor will only assign default values to properties that have it defined,
func NewModifyIsolatedMarginRiskLimtRespWithDefaults() *ModifyIsolatedMarginRiskLimtResp {
	this := ModifyIsolatedMarginRiskLimtResp{}
	return &this
}

func (o *ModifyIsolatedMarginRiskLimtResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize
}

func (o *ModifyIsolatedMarginRiskLimtResp) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.Data)
	return err
}

func (o *ModifyIsolatedMarginRiskLimtResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
