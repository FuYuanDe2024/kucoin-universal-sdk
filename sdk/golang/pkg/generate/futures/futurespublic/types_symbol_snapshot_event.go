// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package futurespublic

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// SymbolSnapshotEvent struct for SymbolSnapshotEvent
type SymbolSnapshotEvent struct {
	// common response
	CommonResponse     *types.WsMessage
	HighPrice          float32 `json:"highPrice,omitempty"`
	LastPrice          float32 `json:"lastPrice,omitempty"`
	LowPrice           float32 `json:"lowPrice,omitempty"`
	Price24HoursBefore float32 `json:"price24HoursBefore,omitempty"`
	PriceChg           float32 `json:"priceChg,omitempty"`
	PriceChgPct        float32 `json:"priceChgPct,omitempty"`
	Symbol             string  `json:"symbol,omitempty"`
	Ts                 int64   `json:"ts,omitempty"`
	Turnover           float32 `json:"turnover,omitempty"`
	Volume             float32 `json:"volume,omitempty"`
}

// NewSymbolSnapshotEvent instantiates a new SymbolSnapshotEvent object
// This constructor will assign default values to properties that have it defined
func NewSymbolSnapshotEvent(highPrice float32, lastPrice float32, lowPrice float32, price24HoursBefore float32, priceChg float32, priceChgPct float32, symbol string, ts int64, turnover float32, volume float32) *SymbolSnapshotEvent {
	this := SymbolSnapshotEvent{}
	this.HighPrice = highPrice
	this.LastPrice = lastPrice
	this.LowPrice = lowPrice
	this.Price24HoursBefore = price24HoursBefore
	this.PriceChg = priceChg
	this.PriceChgPct = priceChgPct
	this.Symbol = symbol
	this.Ts = ts
	this.Turnover = turnover
	this.Volume = volume
	return &this
}

// NewSymbolSnapshotEventWithDefaults instantiates a new SymbolSnapshotEvent object
// This constructor will only assign default values to properties that have it defined,
func NewSymbolSnapshotEventWithDefaults() *SymbolSnapshotEvent {
	this := SymbolSnapshotEvent{}
	return &this
}

func (o *SymbolSnapshotEvent) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["highPrice"] = o.HighPrice
	toSerialize["lastPrice"] = o.LastPrice
	toSerialize["lowPrice"] = o.LowPrice
	toSerialize["price24HoursBefore"] = o.Price24HoursBefore
	toSerialize["priceChg"] = o.PriceChg
	toSerialize["priceChgPct"] = o.PriceChgPct
	toSerialize["symbol"] = o.Symbol
	toSerialize["ts"] = o.Ts
	toSerialize["turnover"] = o.Turnover
	toSerialize["volume"] = o.Volume
	return toSerialize
}

type SymbolSnapshotEventCallback func(topic string, subject string, data *SymbolSnapshotEvent) error

type SymbolSnapshotEventCallbackWrapper struct {
	callback SymbolSnapshotEventCallback
}

func (w *SymbolSnapshotEventCallbackWrapper) OnMessage(msg *types.WsMessage) error {
	obj := &SymbolSnapshotEvent{}
	err := json.Unmarshal(msg.RawData, obj)
	if err != nil {
		return err
	}
	obj.CommonResponse = msg
	return w.callback(msg.Topic, msg.Subject, obj)
}
