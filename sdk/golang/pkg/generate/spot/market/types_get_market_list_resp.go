// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package market

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetMarketListResp struct for GetMarketListResp
type GetMarketListResp struct {
	// common response
	CommonResponse *types.RestResponse
	Data           []string `json:"data,omitempty"`
}

// NewGetMarketListResp instantiates a new GetMarketListResp object
// This constructor will assign default values to properties that have it defined
func NewGetMarketListResp(data []string) *GetMarketListResp {
	this := GetMarketListResp{}
	this.Data = data
	return &this
}

// NewGetMarketListRespWithDefaults instantiates a new GetMarketListResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetMarketListRespWithDefaults() *GetMarketListResp {
	this := GetMarketListResp{}
	return &this
}

func (o *GetMarketListResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize
}

func (o *GetMarketListResp) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.Data)
	return err
}

func (o *GetMarketListResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
