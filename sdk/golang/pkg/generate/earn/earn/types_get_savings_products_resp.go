// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package earn

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetSavingsProductsResp struct for GetSavingsProductsResp
type GetSavingsProductsResp struct {
	// common response
	CommonResponse *types.RestResponse
	Data           []GetSavingsProductsData `json:"data,omitempty"`
}

// NewGetSavingsProductsResp instantiates a new GetSavingsProductsResp object
// This constructor will assign default values to properties that have it defined
func NewGetSavingsProductsResp(data []GetSavingsProductsData) *GetSavingsProductsResp {
	this := GetSavingsProductsResp{}
	this.Data = data
	return &this
}

// NewGetSavingsProductsRespWithDefaults instantiates a new GetSavingsProductsResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetSavingsProductsRespWithDefaults() *GetSavingsProductsResp {
	this := GetSavingsProductsResp{}
	return &this
}

func (o *GetSavingsProductsResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize
}

func (o *GetSavingsProductsResp) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.Data)
	return err
}

func (o *GetSavingsProductsResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
