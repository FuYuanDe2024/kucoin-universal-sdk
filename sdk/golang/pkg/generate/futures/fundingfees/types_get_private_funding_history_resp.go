// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package fundingfees

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetPrivateFundingHistoryResp struct for GetPrivateFundingHistoryResp
type GetPrivateFundingHistoryResp struct {
	// common response
	CommonResponse *types.RestResponse
	DataList       []GetPrivateFundingHistoryDataList `json:"dataList,omitempty"`
	// Whether there are more pages
	HasMore bool `json:"hasMore,omitempty"`
}

// NewGetPrivateFundingHistoryResp instantiates a new GetPrivateFundingHistoryResp object
// This constructor will assign default values to properties that have it defined
func NewGetPrivateFundingHistoryResp(dataList []GetPrivateFundingHistoryDataList, hasMore bool) *GetPrivateFundingHistoryResp {
	this := GetPrivateFundingHistoryResp{}
	this.DataList = dataList
	this.HasMore = hasMore
	return &this
}

// NewGetPrivateFundingHistoryRespWithDefaults instantiates a new GetPrivateFundingHistoryResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetPrivateFundingHistoryRespWithDefaults() *GetPrivateFundingHistoryResp {
	this := GetPrivateFundingHistoryResp{}
	return &this
}

func (o *GetPrivateFundingHistoryResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["dataList"] = o.DataList
	toSerialize["hasMore"] = o.HasMore
	return toSerialize
}

func (o *GetPrivateFundingHistoryResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
