// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package withdrawal

// CancelWithdrawalReq struct for CancelWithdrawalReq
type CancelWithdrawalReq struct {
	// Path parameter, a unique ID for a withdrawalId
	WithdrawalId *string `json:"withdrawalId,omitempty" path:"withdrawalId" url:"-"`
}

// NewCancelWithdrawalReq instantiates a new CancelWithdrawalReq object
// This constructor will assign default values to properties that have it defined
func NewCancelWithdrawalReq() *CancelWithdrawalReq {
	this := CancelWithdrawalReq{}
	return &this
}

// NewCancelWithdrawalReqWithDefaults instantiates a new CancelWithdrawalReq object
// This constructor will only assign default values to properties that have it defined,
func NewCancelWithdrawalReqWithDefaults() *CancelWithdrawalReq {
	this := CancelWithdrawalReq{}
	return &this
}

func (o *CancelWithdrawalReq) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["withdrawalId"] = o.WithdrawalId
	return toSerialize
}

type CancelWithdrawalReqBuilder struct {
	obj *CancelWithdrawalReq
}

func NewCancelWithdrawalReqBuilder() *CancelWithdrawalReqBuilder {
	return &CancelWithdrawalReqBuilder{obj: NewCancelWithdrawalReqWithDefaults()}
}

// Path parameter, a unique ID for a withdrawalId
func (builder *CancelWithdrawalReqBuilder) SetWithdrawalId(value string) *CancelWithdrawalReqBuilder {
	builder.obj.WithdrawalId = &value
	return builder
}

func (builder *CancelWithdrawalReqBuilder) Build() *CancelWithdrawalReq {
	return builder.obj
}
