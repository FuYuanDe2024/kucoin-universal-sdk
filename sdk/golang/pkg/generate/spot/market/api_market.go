// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package market

import (
	"context"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/internal/interfaces"
)

type MarketAPI interface {

	// GetAnnouncements Get Announcements
	// Description: This interface can obtain the latest news announcements, and the default page search is for announcements within a month.
	// Documentation: https://www.kucoin.com/docs-new/api-3470157
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 20     |
	// +---------------------+--------+
	GetAnnouncements(req *GetAnnouncementsReq, ctx context.Context) (*GetAnnouncementsResp, error)

	// GetCurrency Get Currency
	// Description: Request via this endpoint to get the currency details of a specified currency
	// Documentation: https://www.kucoin.com/docs-new/api-3470155
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetCurrency(req *GetCurrencyReq, ctx context.Context) (*GetCurrencyResp, error)

	// GetAllCurrencies Get All Currencies
	// Description: Request via this endpoint to get the currency list.Not all currencies currently can be used for trading.
	// Documentation: https://www.kucoin.com/docs-new/api-3470152
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetAllCurrencies(ctx context.Context) (*GetAllCurrenciesResp, error)

	// GetSymbol Get Symbol
	// Description: Request via this endpoint to get detail currency pairs for trading. If you want to get the market information of the trading symbol, please use Get All Tickers.
	// Documentation: https://www.kucoin.com/docs-new/api-3470159
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 4      |
	// +---------------------+--------+
	GetSymbol(req *GetSymbolReq, ctx context.Context) (*GetSymbolResp, error)

	// GetAllSymbols Get All Symbols
	// Description: Request via this endpoint to get a list of available currency pairs for trading. If you want to get the market information of the trading symbol, please use Get All Tickers.
	// Documentation: https://www.kucoin.com/docs-new/api-3470154
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 4      |
	// +---------------------+--------+
	GetAllSymbols(req *GetAllSymbolsReq, ctx context.Context) (*GetAllSymbolsResp, error)

	// GetTicker Get Ticker
	// Description: Request via this endpoint to get Level 1 Market Data. The returned value includes the best bid price and size, the best ask price and size as well as the last traded price and the last traded size.
	// Documentation: https://www.kucoin.com/docs-new/api-3470160
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 2      |
	// +---------------------+--------+
	GetTicker(req *GetTickerReq, ctx context.Context) (*GetTickerResp, error)

	// GetAllTickers Get All Tickers
	// Description: Request market tickers for all the trading pairs in the market (including 24h volume), takes a snapshot every 2 seconds.  On the rare occasion that we will change the currency name, if you still want the changed symbol name, you can use the symbolName field instead of the symbol field via “Get all tickers” endpoint.
	// Documentation: https://www.kucoin.com/docs-new/api-3470167
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 15     |
	// +---------------------+--------+
	GetAllTickers(ctx context.Context) (*GetAllTickersResp, error)

	// GetTradeHistory Get Trade History
	// Description: Request via this endpoint to get the trade history of the specified symbol, the returned quantity is the last 100 transaction records.
	// Documentation: https://www.kucoin.com/docs-new/api-3470162
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetTradeHistory(req *GetTradeHistoryReq, ctx context.Context) (*GetTradeHistoryResp, error)

	// GetKlines Get Klines
	// Description: Get the Kline of the symbol. Data are returned in grouped buckets based on requested type. For each query, the system would return at most 1500 pieces of data. To obtain more data, please page the data by time.
	// Documentation: https://www.kucoin.com/docs-new/api-3470163
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetKlines(req *GetKlinesReq, ctx context.Context) (*GetKlinesResp, error)

	// GetPartOrderBook Get Part OrderBook
	// Description: Query for part orderbook depth data. (aggregated by price)  You are recommended to request via this endpoint as the system reponse would be faster and cosume less traffic.
	// Documentation: https://www.kucoin.com/docs-new/api-3470165
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 2      |
	// +---------------------+--------+
	GetPartOrderBook(req *GetPartOrderBookReq, ctx context.Context) (*GetPartOrderBookResp, error)

	// GetFullOrderBook Get Full OrderBook
	// Description: Query for Full orderbook depth data. (aggregated by price)  It is generally used by professional traders because it uses more server resources and traffic, and we have strict access rate limit control.  To maintain up-to-date Order Book, please use Websocket incremental feed after retrieving the OrderBook.
	// Documentation: https://www.kucoin.com/docs-new/api-3470164
	// +---------------------+---------+
	// | Extra API Info      | Value   |
	// +---------------------+---------+
	// | API-DOMAIN          | SPOT    |
	// | API-CHANNEL         | PRIVATE |
	// | API-PERMISSION      | GENERAL |
	// | API-RATE-LIMIT-POOL | SPOT    |
	// | API-RATE-LIMIT      | 3       |
	// +---------------------+---------+
	GetFullOrderBook(req *GetFullOrderBookReq, ctx context.Context) (*GetFullOrderBookResp, error)

	// GetFiatPrice Get Fiat Price
	// Description: Request via this endpoint to get the fiat price of the currencies for the available trading pairs.
	// Documentation: https://www.kucoin.com/docs-new/api-3470153
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetFiatPrice(req *GetFiatPriceReq, ctx context.Context) (*GetFiatPriceResp, error)

	// Get24hrStats Get 24hr Stats
	// Description: Request via this endpoint to get the statistics of the specified ticker in the last 24 hours.
	// Documentation: https://www.kucoin.com/docs-new/api-3470161
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 15     |
	// +---------------------+--------+
	Get24hrStats(req *Get24hrStatsReq, ctx context.Context) (*Get24hrStatsResp, error)

	// GetMarketList Get Market List
	// Description: Request via this endpoint to get the transaction currency for the entire trading market.
	// Documentation: https://www.kucoin.com/docs-new/api-3470166
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetMarketList(ctx context.Context) (*GetMarketListResp, error)

	// GetServerTime Get Server Time
	// Description: Get the server time.
	// Documentation: https://www.kucoin.com/docs-new/api-3470156
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetServerTime(ctx context.Context) (*GetServerTimeResp, error)

	// GetServiceStatus Get Service Status
	// Description: Get the service status
	// Documentation: https://www.kucoin.com/docs-new/api-3470158
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 3      |
	// +---------------------+--------+
	GetServiceStatus(ctx context.Context) (*GetServiceStatusResp, error)

	// GetPublicToken Get Public Token - Spot/Margin
	// Description: This interface can obtain the token required for websocket to establish a Spot/Margin connection. If you need use public channels (e.g. all public market data), please make request as follows to obtain the server list and public token
	// Documentation: https://www.kucoin.com/docs-new/api-3470294
	// +---------------------+--------+
	// | Extra API Info      | Value  |
	// +---------------------+--------+
	// | API-DOMAIN          | SPOT   |
	// | API-CHANNEL         | PUBLIC |
	// | API-PERMISSION      | NULL   |
	// | API-RATE-LIMIT-POOL | PUBLIC |
	// | API-RATE-LIMIT      | 10     |
	// +---------------------+--------+
	GetPublicToken(ctx context.Context) (*GetPublicTokenResp, error)

	// GetPrivateToken Get Private Token - Spot/Margin
	// Description: This interface can obtain the token required for websocket to establish a Spot/Margin private connection. If you need use private channels(e.g. account balance notice), please make request as follows to obtain the server list and private token
	// Documentation: https://www.kucoin.com/docs-new/api-3470295
	// +---------------------+---------+
	// | Extra API Info      | Value   |
	// +---------------------+---------+
	// | API-DOMAIN          | SPOT    |
	// | API-CHANNEL         | PRIVATE |
	// | API-PERMISSION      | GENERAL |
	// | API-RATE-LIMIT-POOL | SPOT    |
	// | API-RATE-LIMIT      | 10      |
	// +---------------------+---------+
	GetPrivateToken(ctx context.Context) (*GetPrivateTokenResp, error)
}

type MarketAPIImpl struct {
	transport interfaces.Transport
}

func NewMarketAPIImp(transport interfaces.Transport) *MarketAPIImpl {
	return &MarketAPIImpl{transport: transport}
}

func (impl *MarketAPIImpl) GetAnnouncements(req *GetAnnouncementsReq, ctx context.Context) (*GetAnnouncementsResp, error) {
	resp := &GetAnnouncementsResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/announcements", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetCurrency(req *GetCurrencyReq, ctx context.Context) (*GetCurrencyResp, error) {
	resp := &GetCurrencyResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/currencies/{currency}", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetAllCurrencies(ctx context.Context) (*GetAllCurrenciesResp, error) {
	resp := &GetAllCurrenciesResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/currencies", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetSymbol(req *GetSymbolReq, ctx context.Context) (*GetSymbolResp, error) {
	resp := &GetSymbolResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v2/symbols/{symbol}", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetAllSymbols(req *GetAllSymbolsReq, ctx context.Context) (*GetAllSymbolsResp, error) {
	resp := &GetAllSymbolsResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v2/symbols", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetTicker(req *GetTickerReq, ctx context.Context) (*GetTickerResp, error) {
	resp := &GetTickerResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/orderbook/level1", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetAllTickers(ctx context.Context) (*GetAllTickersResp, error) {
	resp := &GetAllTickersResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/allTickers", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetTradeHistory(req *GetTradeHistoryReq, ctx context.Context) (*GetTradeHistoryResp, error) {
	resp := &GetTradeHistoryResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/histories", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetKlines(req *GetKlinesReq, ctx context.Context) (*GetKlinesResp, error) {
	resp := &GetKlinesResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/candles", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetPartOrderBook(req *GetPartOrderBookReq, ctx context.Context) (*GetPartOrderBookResp, error) {
	resp := &GetPartOrderBookResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/orderbook/level2_{size}", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetFullOrderBook(req *GetFullOrderBookReq, ctx context.Context) (*GetFullOrderBookResp, error) {
	resp := &GetFullOrderBookResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v3/market/orderbook/level2", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetFiatPrice(req *GetFiatPriceReq, ctx context.Context) (*GetFiatPriceResp, error) {
	resp := &GetFiatPriceResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/prices", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) Get24hrStats(req *Get24hrStatsReq, ctx context.Context) (*Get24hrStatsResp, error) {
	resp := &Get24hrStatsResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/market/stats", req, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetMarketList(ctx context.Context) (*GetMarketListResp, error) {
	resp := &GetMarketListResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/markets", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetServerTime(ctx context.Context) (*GetServerTimeResp, error) {
	resp := &GetServerTimeResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/timestamp", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetServiceStatus(ctx context.Context) (*GetServiceStatusResp, error) {
	resp := &GetServiceStatusResp{}
	err := impl.transport.Call(ctx, "spot", false, "Get", "/api/v1/status", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetPublicToken(ctx context.Context) (*GetPublicTokenResp, error) {
	resp := &GetPublicTokenResp{}
	err := impl.transport.Call(ctx, "spot", false, "Post", "/api/v1/bullet-public", nil, resp, false)
	return resp, err
}

func (impl *MarketAPIImpl) GetPrivateToken(ctx context.Context) (*GetPrivateTokenResp, error) {
	resp := &GetPrivateTokenResp{}
	err := impl.transport.Call(ctx, "spot", false, "Post", "/api/v1/bullet-private", nil, resp, false)
	return resp, err
}
