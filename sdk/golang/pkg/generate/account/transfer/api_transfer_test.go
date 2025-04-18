package transfer

import (
	"encoding/json"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransferGetTransferQuotasReqModel(t *testing.T) {
	// GetTransferQuotas
	// Get Transfer Quotas
	// /api/v1/accounts/transferable

	data := "{\"currency\": \"BTC\", \"type\": \"MAIN\", \"tag\": \"ETH-USDT\"}"
	req := &GetTransferQuotasReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferGetTransferQuotasRespModel(t *testing.T) {
	// GetTransferQuotas
	// Get Transfer Quotas
	// /api/v1/accounts/transferable

	data := "{\"code\":\"200000\",\"data\":{\"currency\":\"USDT\",\"balance\":\"10.5\",\"available\":\"10.5\",\"holds\":\"0\",\"transferable\":\"10.5\"}}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetTransferQuotasResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferFlexTransferReqModel(t *testing.T) {
	// FlexTransfer
	// Flex Transfer
	// /api/v3/accounts/universal-transfer

	data := "{\"clientOid\": \"64ccc0f164781800010d8c09\", \"type\": \"PARENT_TO_SUB\", \"currency\": \"USDT\", \"amount\": \"0.01\", \"fromAccountType\": \"TRADE\", \"toUserId\": \"63743f07e0c5230001761d08\", \"toAccountType\": \"TRADE\"}"
	req := &FlexTransferReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferFlexTransferRespModel(t *testing.T) {
	// FlexTransfer
	// Flex Transfer
	// /api/v3/accounts/universal-transfer

	data := "{\n    \"code\": \"200000\",\n    \"data\": {\n        \"orderId\": \"6705f7248c6954000733ecac\"\n    }\n}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &FlexTransferResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferSubAccountTransferReqModel(t *testing.T) {
	// SubAccountTransfer
	// SubAccount Transfer
	// /api/v2/accounts/sub-transfer

	data := "{\"clientOid\": \"64ccc0f164781800010d8c09\", \"currency\": \"USDT\", \"amount\": \"0.01\", \"direction\": \"OUT\", \"accountType\": \"MAIN\", \"subAccountType\": \"MAIN\", \"subUserId\": \"63743f07e0c5230001761d08\"}"
	req := &SubAccountTransferReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferSubAccountTransferRespModel(t *testing.T) {
	// SubAccountTransfer
	// SubAccount Transfer
	// /api/v2/accounts/sub-transfer

	data := "{\"code\":\"200000\",\"data\":{\"orderId\":\"670be6b0b1b9080007040a9b\"}}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &SubAccountTransferResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferInnerTransferReqModel(t *testing.T) {
	// InnerTransfer
	// Inner Transfer
	// /api/v2/accounts/inner-transfer

	data := "{\"clientOid\": \"64ccc0f164781800010d8c09\", \"currency\": \"USDT\", \"amount\": \"0.01\", \"from\": \"main\", \"to\": \"trade\"}"
	req := &InnerTransferReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferInnerTransferRespModel(t *testing.T) {
	// InnerTransfer
	// Inner Transfer
	// /api/v2/accounts/inner-transfer

	data := "{\"code\":\"200000\",\"data\":{\"orderId\":\"670beb3482a1bb0007dec644\"}}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &InnerTransferResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferFuturesAccountTransferOutReqModel(t *testing.T) {
	// FuturesAccountTransferOut
	// Futures Account Transfer Out
	// /api/v3/transfer-out

	data := "{\"currency\": \"USDT\", \"amount\": 0.01, \"recAccountType\": \"MAIN\"}"
	req := &FuturesAccountTransferOutReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferFuturesAccountTransferOutRespModel(t *testing.T) {
	// FuturesAccountTransferOut
	// Futures Account Transfer Out
	// /api/v3/transfer-out

	data := "{\n    \"code\": \"200000\",\n    \"data\": {\n        \"applyId\": \"670bf84c577f6c00017a1c48\",\n        \"bizNo\": \"670bf84c577f6c00017a1c47\",\n        \"payAccountType\": \"CONTRACT\",\n        \"payTag\": \"DEFAULT\",\n        \"remark\": \"\",\n        \"recAccountType\": \"MAIN\",\n        \"recTag\": \"DEFAULT\",\n        \"recRemark\": \"\",\n        \"recSystem\": \"KUCOIN\",\n        \"status\": \"PROCESSING\",\n        \"currency\": \"USDT\",\n        \"amount\": \"0.01\",\n        \"fee\": \"0\",\n        \"sn\": 1519769124134806,\n        \"reason\": \"\",\n        \"createdAt\": 1728837708000,\n        \"updatedAt\": 1728837708000\n    }\n}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &FuturesAccountTransferOutResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferFuturesAccountTransferInReqModel(t *testing.T) {
	// FuturesAccountTransferIn
	// Futures Account Transfer In
	// /api/v1/transfer-in

	data := "{\"currency\": \"USDT\", \"amount\": 0.01, \"payAccountType\": \"MAIN\"}"
	req := &FuturesAccountTransferInReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferFuturesAccountTransferInRespModel(t *testing.T) {
	// FuturesAccountTransferIn
	// Futures Account Transfer In
	// /api/v1/transfer-in

	data := "{\"code\":\"200000\",\"data\":null}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &FuturesAccountTransferInResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestTransferGetFuturesAccountTransferOutLedgerReqModel(t *testing.T) {
	// GetFuturesAccountTransferOutLedger
	// Get Futures Account Transfer Out Ledger
	// /api/v1/transfer-list

	data := "{\"currency\": \"XBT\", \"type\": \"MAIN\", \"tag\": [\"mock_a\", \"mock_b\"], \"startAt\": 1728663338000, \"endAt\": 1728692138000, \"currentPage\": 1, \"pageSize\": 50}"
	req := &GetFuturesAccountTransferOutLedgerReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestTransferGetFuturesAccountTransferOutLedgerRespModel(t *testing.T) {
	// GetFuturesAccountTransferOutLedger
	// Get Futures Account Transfer Out Ledger
	// /api/v1/transfer-list

	data := "{\"code\":\"200000\",\"data\":{\"currentPage\":1,\"pageSize\":50,\"totalNum\":1,\"totalPage\":1,\"items\":[{\"applyId\":\"670bf84c577f6c00017a1c48\",\"currency\":\"USDT\",\"recRemark\":\"\",\"recSystem\":\"KUCOIN\",\"status\":\"SUCCESS\",\"amount\":\"0.01\",\"reason\":\"\",\"offset\":1519769124134806,\"createdAt\":1728837708000,\"remark\":\"\"}]}}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetFuturesAccountTransferOutLedgerResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}
