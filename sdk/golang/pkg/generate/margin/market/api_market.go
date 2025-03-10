// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package market

import (
	"context"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/internal/interfaces"
)

type MarketAPI interface {

	// GetCrossMarginSymbols Get Symbols - Cross Margin
	// Description: This endpoint allows querying the configuration of cross margin symbol.
	// Documentation: https://www.kucoin.com/docs-new/api-3470189
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetCrossMarginSymbols(req *GetCrossMarginSymbolsReq, ctx context.Context) (*GetCrossMarginSymbolsResp, error)

	// GetMarginConfig Get Margin Config
	// Description: Request via this endpoint to get the configure info of the cross margin.
	// Documentation: https://www.kucoin.com/docs-new/api-3470190
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | SPOT   |
	// | API-RATE-LIMIT      | 25     |
	// +---------------------+--------+
	GetMarginConfig(ctx context.Context) (*GetMarginConfigResp, error)

	// GetETFInfo Get ETF Info
	// Description: This interface returns leveraged token information
	// Documentation: https://www.kucoin.com/docs-new/api-3470191
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetETFInfo(req *GetETFInfoReq, ctx context.Context) (*GetETFInfoResp, error)

	// GetMarkPriceList Get Mark Price List
	// Description: This endpoint returns the current Mark price for all margin trading pairs.
	// Documentation: https://www.kucoin.com/docs-new/api-3470192
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 10     |
	// +---------------------+--------+
	GetMarkPriceList(ctx context.Context) (*GetMarkPriceListResp, error)

	// GetMarkPriceDetail Get Mark Price Detail
	// Description: This endpoint returns the current Mark price for specified margin trading pairs.
	// Documentation: https://www.kucoin.com/docs-new/api-3470193
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 2      |
	// +---------------------+--------+
	GetMarkPriceDetail(req *GetMarkPriceDetailReq, ctx context.Context) (*GetMarkPriceDetailResp, error)

	// GetIsolatedMarginSymbols Get Symbols - Isolated Margin
	// Description: This endpoint allows querying the configuration of isolated margin symbol.
	// Documentation: https://www.kucoin.com/docs-new/api-3470194
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetIsolatedMarginSymbols(ctx context.Context) (*GetIsolatedMarginSymbolsResp, error)
}

type MarketAPIImpl struct {
	transport interfaces.Transport
}

func NewMarketAPIImp(transport interfaces.Transport) *MarketAPIImpl {
	return &MarketAPIImpl{transport: transport}
}

func (impl *MarketAPIImpl) GetCrossMarginSymbols(req *GetCrossMarginSymbolsReq, ctx context.Context) (*GetCrossMarginSymbolsResp, error) {
	resp := &GetCrossMarginSymbolsResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/margin/symbols", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetMarginConfig(ctx context.Context) (*GetMarginConfigResp, error) {
	resp := &GetMarginConfigResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/margin/config", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetETFInfo(req *GetETFInfoReq, ctx context.Context) (*GetETFInfoResp, error) {
	resp := &GetETFInfoResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/etf/info", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetMarkPriceList(ctx context.Context) (*GetMarkPriceListResp, error) {
	resp := &GetMarkPriceListResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/mark-price/all-symbols", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetMarkPriceDetail(req *GetMarkPriceDetailReq, ctx context.Context) (*GetMarkPriceDetailResp, error) {
	resp := &GetMarkPriceDetailResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/mark-price/{symbol}/current", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetIsolatedMarginSymbols(ctx context.Context) (*GetIsolatedMarginSymbolsResp, error) {
	resp := &GetIsolatedMarginSymbolsResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/isolated/symbols", nil, resp, false)
	return resp, err
}
