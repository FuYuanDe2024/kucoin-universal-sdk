// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package spotpublic

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// SymbolSnapshotEvent struct for SymbolSnapshotEvent
type SymbolSnapshotEvent struct {
	// common response
	CommonResponse *types.WsMessage
	Sequence       string             `json:"sequence,omitempty"`
	Data           SymbolSnapshotData `json:"data,omitempty"`
}

// NewSymbolSnapshotEvent instantiates a new SymbolSnapshotEvent object
// This constructor will assign default values to properties that have it defined
func NewSymbolSnapshotEvent(sequence string, data SymbolSnapshotData) *SymbolSnapshotEvent {
	this := SymbolSnapshotEvent{}
	this.Sequence = sequence
	this.Data = data
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
	toSerialize["sequence"] = o.Sequence
	toSerialize["data"] = o.Data
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
