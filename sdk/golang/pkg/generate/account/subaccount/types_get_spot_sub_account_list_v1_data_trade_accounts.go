// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package subaccount

// GetSpotSubAccountListV1DataTradeAccounts struct for GetSpotSubAccountListV1DataTradeAccounts
type GetSpotSubAccountListV1DataTradeAccounts struct {
	// The currency of the account.
	Currency *string `json:"currency,omitempty"`
	// Total funds in the account.
	Balance *string `json:"balance,omitempty"`
	// Funds available to withdraw or trade.
	Available *string `json:"available,omitempty"`
	// Funds on hold (not available for use).
	Holds *string `json:"holds,omitempty"`
	// Calculated on this currency.
	BaseCurrency *string `json:"baseCurrency,omitempty"`
	// The base currency price.
	BaseCurrencyPrice *string `json:"baseCurrencyPrice,omitempty"`
	// The base currency amount.
	BaseAmount *string `json:"baseAmount,omitempty"`
	Tag        *string `json:"tag,omitempty"`
}

// NewGetSpotSubAccountListV1DataTradeAccounts instantiates a new GetSpotSubAccountListV1DataTradeAccounts object
// This constructor will assign default values to properties that have it defined
func NewGetSpotSubAccountListV1DataTradeAccounts() *GetSpotSubAccountListV1DataTradeAccounts {
	this := GetSpotSubAccountListV1DataTradeAccounts{}
	return &this
}

// NewGetSpotSubAccountListV1DataTradeAccountsWithDefaults instantiates a new GetSpotSubAccountListV1DataTradeAccounts object
// This constructor will only assign default values to properties that have it defined,
func NewGetSpotSubAccountListV1DataTradeAccountsWithDefaults() *GetSpotSubAccountListV1DataTradeAccounts {
	this := GetSpotSubAccountListV1DataTradeAccounts{}
	return &this
}

func (o *GetSpotSubAccountListV1DataTradeAccounts) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["balance"] = o.Balance
	toSerialize["available"] = o.Available
	toSerialize["holds"] = o.Holds
	toSerialize["baseCurrency"] = o.BaseCurrency
	toSerialize["baseCurrencyPrice"] = o.BaseCurrencyPrice
	toSerialize["baseAmount"] = o.BaseAmount
	toSerialize["tag"] = o.Tag
	return toSerialize
}
