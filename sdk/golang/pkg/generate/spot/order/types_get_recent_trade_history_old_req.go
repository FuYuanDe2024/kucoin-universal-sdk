// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

// GetRecentTradeHistoryOldReq struct for GetRecentTradeHistoryOldReq
type GetRecentTradeHistoryOldReq struct {
	// Current request page.
	CurrentPage *int32 `json:"currentPage,omitempty" url:"currentPage,omitempty"`
	// Number of results per request. Minimum is 10, maximum is 500.
	PageSize *int32 `json:"pageSize,omitempty" url:"pageSize,omitempty"`
}

// NewGetRecentTradeHistoryOldReq instantiates a new GetRecentTradeHistoryOldReq object
// This constructor will assign default values to properties that have it defined
func NewGetRecentTradeHistoryOldReq() *GetRecentTradeHistoryOldReq {
	this := GetRecentTradeHistoryOldReq{}
	var currentPage int32 = 1
	this.CurrentPage = &currentPage
	return &this
}

// NewGetRecentTradeHistoryOldReqWithDefaults instantiates a new GetRecentTradeHistoryOldReq object
// This constructor will only assign default values to properties that have it defined,
func NewGetRecentTradeHistoryOldReqWithDefaults() *GetRecentTradeHistoryOldReq {
	this := GetRecentTradeHistoryOldReq{}
	var currentPage int32 = 1
	this.CurrentPage = &currentPage
	return &this
}

func (o *GetRecentTradeHistoryOldReq) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["currentPage"] = o.CurrentPage
	toSerialize["pageSize"] = o.PageSize
	return toSerialize
}

type GetRecentTradeHistoryOldReqBuilder struct {
	obj *GetRecentTradeHistoryOldReq
}

func NewGetRecentTradeHistoryOldReqBuilder() *GetRecentTradeHistoryOldReqBuilder {
	return &GetRecentTradeHistoryOldReqBuilder{obj: NewGetRecentTradeHistoryOldReqWithDefaults()}
}

// Current request page.
func (builder *GetRecentTradeHistoryOldReqBuilder) SetCurrentPage(value int32) *GetRecentTradeHistoryOldReqBuilder {
	builder.obj.CurrentPage = &value
	return builder
}

// Number of results per request. Minimum is 10, maximum is 500.
func (builder *GetRecentTradeHistoryOldReqBuilder) SetPageSize(value int32) *GetRecentTradeHistoryOldReqBuilder {
	builder.obj.PageSize = &value
	return builder
}

func (builder *GetRecentTradeHistoryOldReqBuilder) Build() *GetRecentTradeHistoryOldReq {
	return builder.obj
}
