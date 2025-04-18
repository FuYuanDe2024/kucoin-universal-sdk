// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package futures

// CancelOrderByClientOidReq struct for CancelOrderByClientOidReq
type CancelOrderByClientOidReq struct {
	// Symbol of the contract, Please refer to [Get Symbol endpoint: symbol](https://www.kucoin.com/docs-new/api-3470220)
	Symbol *string `json:"symbol,omitempty" url:"symbol,omitempty"`
	// The user self-defined order id.
	ClientOid *string `json:"clientOid,omitempty" url:"clientOid,omitempty"`
}

// NewCancelOrderByClientOidReq instantiates a new CancelOrderByClientOidReq object
// This constructor will assign default values to properties that have it defined
func NewCancelOrderByClientOidReq() *CancelOrderByClientOidReq {
	this := CancelOrderByClientOidReq{}
	return &this
}

// NewCancelOrderByClientOidReqWithDefaults instantiates a new CancelOrderByClientOidReq object
// This constructor will only assign default values to properties that have it defined,
func NewCancelOrderByClientOidReqWithDefaults() *CancelOrderByClientOidReq {
	this := CancelOrderByClientOidReq{}
	return &this
}

func (o *CancelOrderByClientOidReq) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["clientOid"] = o.ClientOid
	return toSerialize
}

type CancelOrderByClientOidReqBuilder struct {
	obj *CancelOrderByClientOidReq
}

func NewCancelOrderByClientOidReqBuilder() *CancelOrderByClientOidReqBuilder {
	return &CancelOrderByClientOidReqBuilder{obj: NewCancelOrderByClientOidReqWithDefaults()}
}

// Symbol of the contract, Please refer to [Get Symbol endpoint: symbol](https://www.kucoin.com/docs-new/api-3470220)
func (builder *CancelOrderByClientOidReqBuilder) SetSymbol(value string) *CancelOrderByClientOidReqBuilder {
	builder.obj.Symbol = &value
	return builder
}

// The user self-defined order id.
func (builder *CancelOrderByClientOidReqBuilder) SetClientOid(value string) *CancelOrderByClientOidReqBuilder {
	builder.obj.ClientOid = &value
	return builder
}

func (builder *CancelOrderByClientOidReqBuilder) Build() *CancelOrderByClientOidReq {
	return builder.obj
}
