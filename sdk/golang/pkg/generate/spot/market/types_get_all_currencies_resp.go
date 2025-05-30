// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package market

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetAllCurrenciesResp struct for GetAllCurrenciesResp
type GetAllCurrenciesResp struct {
	// common response
	CommonResponse *types.RestResponse
	Data           []GetAllCurrenciesData `json:"data,omitempty"`
}

// NewGetAllCurrenciesResp instantiates a new GetAllCurrenciesResp object
// This constructor will assign default values to properties that have it defined
func NewGetAllCurrenciesResp(data []GetAllCurrenciesData) *GetAllCurrenciesResp {
	this := GetAllCurrenciesResp{}
	this.Data = data
	return &this
}

// NewGetAllCurrenciesRespWithDefaults instantiates a new GetAllCurrenciesResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetAllCurrenciesRespWithDefaults() *GetAllCurrenciesResp {
	this := GetAllCurrenciesResp{}
	return &this
}

func (o *GetAllCurrenciesResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize
}

func (o *GetAllCurrenciesResp) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.Data)
	return err
}

func (o *GetAllCurrenciesResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
