// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package debit

import (
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// GetInterestHistoryResp struct for GetInterestHistoryResp
type GetInterestHistoryResp struct {
	// common response
	CommonResponse *types.RestResponse
	Timestamp      int64 `json:"timestamp,omitempty"`
	// current page
	CurrentPage int32 `json:"currentPage,omitempty"`
	// page size
	PageSize int32 `json:"pageSize,omitempty"`
	// total number
	TotalNum int32 `json:"totalNum,omitempty"`
	// total page
	TotalPage int32                     `json:"totalPage,omitempty"`
	Items     []GetInterestHistoryItems `json:"items,omitempty"`
}

// NewGetInterestHistoryResp instantiates a new GetInterestHistoryResp object
// This constructor will assign default values to properties that have it defined
func NewGetInterestHistoryResp(timestamp int64, currentPage int32, pageSize int32, totalNum int32, totalPage int32, items []GetInterestHistoryItems) *GetInterestHistoryResp {
	this := GetInterestHistoryResp{}
	this.Timestamp = timestamp
	this.CurrentPage = currentPage
	this.PageSize = pageSize
	this.TotalNum = totalNum
	this.TotalPage = totalPage
	this.Items = items
	return &this
}

// NewGetInterestHistoryRespWithDefaults instantiates a new GetInterestHistoryResp object
// This constructor will only assign default values to properties that have it defined,
func NewGetInterestHistoryRespWithDefaults() *GetInterestHistoryResp {
	this := GetInterestHistoryResp{}
	return &this
}

func (o *GetInterestHistoryResp) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["currentPage"] = o.CurrentPage
	toSerialize["pageSize"] = o.PageSize
	toSerialize["totalNum"] = o.TotalNum
	toSerialize["totalPage"] = o.TotalPage
	toSerialize["items"] = o.Items
	return toSerialize
}

func (o *GetInterestHistoryResp) SetCommonResponse(response *types.RestResponse) {
	o.CommonResponse = response
}
