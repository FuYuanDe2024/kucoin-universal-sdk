// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package order

// BatchCancelOrderOldReq struct for BatchCancelOrderOldReq
type BatchCancelOrderOldReq struct {
	// symbol
	Symbol *string `json:"symbol,omitempty" url:"symbol,omitempty"`
	// The type of trading :TRADE(Spot Trading), MARGIN_TRADE(Cross Margin Trading), MARGIN_ISOLATED_TRADE(Isolated Margin Trading), and the default is TRADE to cancel the spot trading orders.
	TradeType *string `json:"tradeType,omitempty" url:"tradeType,omitempty"`
}

// NewBatchCancelOrderOldReq instantiates a new BatchCancelOrderOldReq object
// This constructor will assign default values to properties that have it defined
func NewBatchCancelOrderOldReq() *BatchCancelOrderOldReq {
	this := BatchCancelOrderOldReq{}
	var tradeType string = "TRADE"
	this.TradeType = &tradeType
	return &this
}

// NewBatchCancelOrderOldReqWithDefaults instantiates a new BatchCancelOrderOldReq object
// This constructor will only assign default values to properties that have it defined,
func NewBatchCancelOrderOldReqWithDefaults() *BatchCancelOrderOldReq {
	this := BatchCancelOrderOldReq{}
	var tradeType string = "TRADE"
	this.TradeType = &tradeType
	return &this
}

func (o *BatchCancelOrderOldReq) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["tradeType"] = o.TradeType
	return toSerialize
}

type BatchCancelOrderOldReqBuilder struct {
	obj *BatchCancelOrderOldReq
}

func NewBatchCancelOrderOldReqBuilder() *BatchCancelOrderOldReqBuilder {
	return &BatchCancelOrderOldReqBuilder{obj: NewBatchCancelOrderOldReqWithDefaults()}
}

// symbol
func (builder *BatchCancelOrderOldReqBuilder) SetSymbol(value string) *BatchCancelOrderOldReqBuilder {
	builder.obj.Symbol = &value
	return builder
}

// The type of trading :TRADE(Spot Trading), MARGIN_TRADE(Cross Margin Trading), MARGIN_ISOLATED_TRADE(Isolated Margin Trading), and the default is TRADE to cancel the spot trading orders.
func (builder *BatchCancelOrderOldReqBuilder) SetTradeType(value string) *BatchCancelOrderOldReqBuilder {
	builder.obj.TradeType = &value
	return builder
}

func (builder *BatchCancelOrderOldReqBuilder) Build() *BatchCancelOrderOldReq {
	return builder.obj
}
