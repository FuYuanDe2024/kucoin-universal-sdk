import {
    ClientOptionBuilder,
    GlobalApiEndpoint,
    GlobalBrokerApiEndpoint,
    GlobalFuturesApiEndpoint, Interceptor,
    TransportOptionBuilder,
} from '@model/index';
import {
    GetFullOrderBookReq,
    GetInterestRateIndexReq,
    GetKlinesReq,
    GetMarkPriceReq,
    GetPartOrderBookReq,
    GetPremiumIndexReq,
    GetSpotIndexPriceReq,
    GetSymbolReq,
    GetTickerReq,
    GetTradeHistoryReq,
    MarketAPI,
} from '@src/generate/futures/market';
import { DefaultClient } from '@api/index';

describe('Auto Test', () => {
    let api: MarketAPI;

    beforeAll(() => {
        const key = process.env.API_KEY || '';
        const secret = process.env.API_SECRET || '';
        const passphrase = process.env.API_PASSPHRASE || '';
        let interceptors: Array<Interceptor> = [
            {
                before: {
                    onFulfilled: (config) => {
                        console.log('debug: ', config.url);
                        return config;
                    },
                },
                after: {},
            },
        ];
        // Set specific options, others will fall back to default values
        const httpTransportOption = new TransportOptionBuilder()
            .setKeepAlive(true)
            .setMaxConnsPerHost(10)
            .setMaxIdleConns(10)
            .setInterceptors(interceptors)
            .build();

        // Create a client using the specified options
        const clientOption = new ClientOptionBuilder()
            .setKey(key)
            .setSecret(secret)
            .setPassphrase(passphrase)
            .setSpotEndpoint(GlobalApiEndpoint)
            .setFuturesEndpoint(GlobalFuturesApiEndpoint)
            .setBrokerEndpoint(GlobalBrokerApiEndpoint)
            .setTransportOption(httpTransportOption)
            .build();

        const client = new DefaultClient(clientOption);

        // Get the Restful Service
        const kucoinRestService = client.restService();
        api = kucoinRestService.getFuturesService().getMarketApi();
    });

    test('getSymbol request test', () => {
        /**
         * getSymbol
         * Get Symbol
         * /api/v1/contracts/{symbol}
         */
        let builder = GetSymbolReq.builder();
        builder.setSymbol('XBTUSDM');
        let req = builder.build();
        let resp = api.getSymbol(req);
        return resp.then((result) => {
            expect(result.symbol).toEqual(expect.anything());
            expect(result.rootSymbol).toEqual(expect.anything());
            expect(result.type).toEqual(expect.anything());
            expect(result.firstOpenDate).toEqual(expect.anything());
            expect(result.baseCurrency).toEqual(expect.anything());
            expect(result.quoteCurrency).toEqual(expect.anything());
            expect(result.settleCurrency).toEqual(expect.anything());
            expect(result.maxOrderQty).toEqual(expect.anything());
            expect(result.maxPrice).toEqual(expect.anything());
            expect(result.lotSize).toEqual(expect.anything());
            expect(result.tickSize).toEqual(expect.anything());
            expect(result.indexPriceTickSize).toEqual(expect.anything());
            expect(result.multiplier).toEqual(expect.anything());
            expect(result.initialMargin).toEqual(expect.anything());
            expect(result.maintainMargin).toEqual(expect.anything());
            expect(result.maxRiskLimit).toEqual(expect.anything());
            expect(result.minRiskLimit).toEqual(expect.anything());
            expect(result.riskStep).toEqual(expect.anything());
            expect(result.makerFeeRate).toEqual(expect.anything());
            expect(result.takerFeeRate).toEqual(expect.anything());
            expect(result.takerFixFee).toEqual(expect.anything());
            expect(result.makerFixFee).toEqual(expect.anything());
            expect(result.isDeleverage).toEqual(expect.anything());
            expect(result.isQuanto).toEqual(expect.anything());
            expect(result.isInverse).toEqual(expect.anything());
            expect(result.markMethod).toEqual(expect.anything());
            expect(result.fairMethod).toEqual(expect.anything());
            expect(result.fundingBaseSymbol).toEqual(expect.anything());
            expect(result.fundingQuoteSymbol).toEqual(expect.anything());
            expect(result.fundingRateSymbol).toEqual(expect.anything());
            expect(result.indexSymbol).toEqual(expect.anything());
            expect(result.status).toEqual(expect.anything());
            expect(result.fundingFeeRate).toEqual(expect.anything());
            expect(result.predictedFundingFeeRate).toEqual(expect.anything());
            expect(result.fundingRateGranularity).toEqual(expect.anything());
            expect(result.openInterest).toEqual(expect.anything());
            expect(result.turnoverOf24h).toEqual(expect.anything());
            expect(result.volumeOf24h).toEqual(expect.anything());
            expect(result.markPrice).toEqual(expect.anything());
            expect(result.indexPrice).toEqual(expect.anything());
            expect(result.lastTradePrice).toEqual(expect.anything());
            expect(result.nextFundingRateTime).toEqual(expect.anything());
            expect(result.maxLeverage).toEqual(expect.anything());
            expect(result.sourceExchanges).toEqual(expect.anything());
            expect(result.premiumsSymbol1M).toEqual(expect.anything());
            expect(result.premiumsSymbol8H).toEqual(expect.anything());
            expect(result.fundingBaseSymbol1M).toEqual(expect.anything());
            expect(result.fundingQuoteSymbol1M).toEqual(expect.anything());
            expect(result.lowPrice).toEqual(expect.anything());
            expect(result.highPrice).toEqual(expect.anything());
            expect(result.priceChgPct).toEqual(expect.anything());
            expect(result.priceChg).toEqual(expect.anything());
            expect(result.k).toEqual(expect.anything());
            expect(result.m).toEqual(expect.anything());
            expect(result.f).toEqual(expect.anything());
            expect(result.mmrLimit).toEqual(expect.anything());
            expect(result.mmrLevConstant).toEqual(expect.anything());
            expect(result.supportCross).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getAllSymbols request test', () => {
        /**
         * getAllSymbols
         * Get All Symbols
         * /api/v1/contracts/active
         */
        let resp = api.getAllSymbols();
        return resp.then((result) => {
            expect(result.data).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getTicker request test', () => {
        /**
         * getTicker
         * Get Ticker
         * /api/v1/ticker
         */
        let builder = GetTickerReq.builder();
        builder.setSymbol('XBTUSDTM');
        let req = builder.build();
        let resp = api.getTicker(req);
        return resp.then((result) => {
            expect(result.sequence).toEqual(expect.anything());
            expect(result.symbol).toEqual(expect.anything());
            expect(result.side).toEqual(expect.anything());
            expect(result.size).toEqual(expect.anything());
            expect(result.tradeId).toEqual(expect.anything());
            expect(result.price).toEqual(expect.anything());
            expect(result.bestBidPrice).toEqual(expect.anything());
            expect(result.bestBidSize).toEqual(expect.anything());
            expect(result.bestAskPrice).toEqual(expect.anything());
            expect(result.bestAskSize).toEqual(expect.anything());
            expect(result.ts).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getAllTickers request test', () => {
        /**
         * getAllTickers
         * Get All Tickers
         * /api/v1/allTickers
         */
        let resp = api.getAllTickers();
        return resp.then((result) => {
            result.data.forEach((item) => {
                expect(item.sequence).toEqual(expect.any(Number));
                expect(item.symbol).toEqual(expect.any(String));
                expect(item.side).toEqual(expect.any(String));
                expect(item.size).toEqual(expect.any(Number));
                expect(item.tradeId).toEqual(expect.any(String));
                expect(item.price).toEqual(expect.any(String));
                expect(item.bestBidPrice).toEqual(expect.any(String));
                expect(item.bestBidSize).toEqual(expect.any(Number));
                expect(item.bestAskPrice).toEqual(expect.any(String));
                expect(item.bestAskSize).toEqual(expect.any(Number));
                expect(item.ts).toEqual(expect.any(Number));
            });
            console.log(resp);
        });
    });

    test('getFullOrderBook request test', () => {
        /**
         * getFullOrderBook
         * Get Full OrderBook
         * /api/v1/level2/snapshot
         */
        let builder = GetFullOrderBookReq.builder();
        builder.setSymbol('XBTUSDTM');
        let req = builder.build();
        let resp = api.getFullOrderBook(req);
        return resp.then((result) => {
            expect(result.sequence).toEqual(expect.anything());
            expect(result.symbol).toEqual(expect.anything());
            expect(result.bids).toEqual(expect.anything());
            expect(result.asks).toEqual(expect.anything());
            expect(result.ts).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getPartOrderBook request test', () => {
        /**
         * getPartOrderBook
         * Get Part OrderBook
         * /api/v1/level2/depth{size}
         */
        let builder = GetPartOrderBookReq.builder();
        builder.setSymbol('XBTUSDTM').setSize('100');
        let req = builder.build();
        let resp = api.getPartOrderBook(req);
        return resp.then((result) => {
            expect(result.sequence).toEqual(expect.anything());
            expect(result.symbol).toEqual(expect.anything());
            expect(result.bids).toEqual(expect.anything());
            expect(result.asks).toEqual(expect.anything());
            expect(result.ts).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getTradeHistory request test', () => {
        /**
         * getTradeHistory
         * Get Trade History
         * /api/v1/trade/history
         */
        let builder = GetTradeHistoryReq.builder();
        builder.setSymbol('XBTUSDTM');
        let req = builder.build();
        let resp = api.getTradeHistory(req);
        return resp.then((result) => {
            result.data.forEach((item) => {
                expect(item.sequence).toEqual(expect.any(Number));
                expect(item.contractId).toEqual(expect.any(Number));
                expect(item.tradeId).toEqual(expect.any(String));
                expect(item.makerOrderId).toEqual(expect.any(String));
                expect(item.takerOrderId).toEqual(expect.any(String));
                expect(item.ts).toEqual(expect.any(Number));
                expect(item.size).toEqual(expect.any(Number));
                expect(item.price).toEqual(expect.any(String));
                expect(item.side).toEqual(expect.any(String));
            });
            console.log(resp);
        });
    });

    test('getKlines request test', () => {
        /**
         * getKlines
         * Get Klines
         * /api/v1/kline/query
         */
        let builder = GetKlinesReq.builder();
        builder
            .setSymbol('XBTUSDTM')
            .setGranularity(GetKlinesReq.GranularityEnum._1)
            .setFrom(1738339200000)
            .setTo(1738425600000);
        let req = builder.build();
        let resp = api.getKlines(req);
        return resp.then((result) => {
            expect(result.data).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getMarkPrice request test', () => {
        /**
         * getMarkPrice
         * Get Mark Price
         * /api/v1/mark-price/{symbol}/current
         */
        let builder = GetMarkPriceReq.builder();
        builder.setSymbol('XBTUSDTM');
        let req = builder.build();
        let resp = api.getMarkPrice(req);
        return resp.then((result) => {
            expect(result.symbol).toEqual(expect.anything());
            expect(result.granularity).toEqual(expect.anything());
            expect(result.timePoint).toEqual(expect.anything());
            expect(result.value).toEqual(expect.anything());
            expect(result.indexPrice).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getSpotIndexPrice request test', () => {
        /**
         * getSpotIndexPrice
         * Get Spot Index Price
         * /api/v1/index/query
         */
        let builder = GetSpotIndexPriceReq.builder();
        builder
            .setSymbol('.KXBTUSDT')
            .setStartAt(1738339200000)
            .setEndAt(1738425600000)
            .setMaxCount(10);
        let req = builder.build();
        let resp = api.getSpotIndexPrice(req);
        return resp.then((result) => {
            result.dataList.forEach(item=> {
                expect(item.symbol).toEqual(expect.any(String));
                expect(item.granularity).toEqual(expect.any(Number));
                expect(item.timePoint).toEqual(expect.any(Number));
                expect(item.value).toEqual(expect.any(Number));
            });
            expect(result.hasMore).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getInterestRateIndex request test', () => {
        /**
         * getInterestRateIndex
         * Get Interest Rate Index
         * /api/v1/interest/query
         */
        let builder = GetInterestRateIndexReq.builder();
        builder.setSymbol('.XBTINT').setStartAt(1738339200000).setEndAt(1738425600000);
        let req = builder.build();
        let resp = api.getInterestRateIndex(req);
        return resp.then((result) => {
            result.dataList.forEach(item=> {
                expect(item.symbol).toEqual(expect.any(String));
                expect(item.granularity).toEqual(expect.any(Number));
                expect(item.timePoint).toEqual(expect.any(Number));
                expect(item.value).toEqual(expect.any(Number));
            });
            expect(result.hasMore).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getPremiumIndex request test', () => {
        /**
         * getPremiumIndex
         * Get Premium Index
         * /api/v1/premium/query
         */
        let builder = GetPremiumIndexReq.builder();
        builder.setSymbol('XBTUSDTMPI').setStartAt(1732464000000).setEndAt(1732521600000);
        let req = builder.build();
        let resp = api.getPremiumIndex(req);
        return resp.then((result) => {
            result.dataList.forEach(item=> {
                expect(item.symbol).toEqual(expect.any(String));
                expect(item.granularity).toEqual(expect.any(Number));
                expect(item.timePoint).toEqual(expect.any(Number));
                expect(item.value).toEqual(expect.any(Number));
            });
            expect(result.hasMore).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('get24hrStats request test', () => {
        /**
         * get24hrStats
         * Get 24hr Stats
         * /api/v1/trade-statistics
         */
        let resp = api.get24hrStats();
        return resp.then((result) => {
            expect(result.turnoverOf24h).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getServerTime request test', () => {
        /**
         * getServerTime
         * Get Server Time
         * /api/v1/timestamp
         */
        let resp = api.getServerTime();
        return resp.then((result) => {
            expect(result.data).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getServiceStatus request test', () => {
        /**
         * getServiceStatus
         * Get Service Status
         * /api/v1/status
         */
        let resp = api.getServiceStatus();
        return resp.then((result) => {
            expect(result.msg).toEqual(expect.anything());
            expect(result.status).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getPublicToken request test', () => {
        /**
         * getPublicToken
         * Get Public Token - Futures
         * /api/v1/bullet-public
         */
        let resp = api.getPublicToken();
        return resp.then((result) => {
            expect(result.token).toEqual(expect.anything());
            expect(result.instanceServers).toEqual(expect.anything());
            console.log(resp);
        });
    });

    test('getPrivateToken request test', () => {
        /**
         * getPrivateToken
         * Get Private Token - Futures
         * /api/v1/bullet-private
         */
        let resp = api.getPrivateToken();
        return resp.then((result) => {
            expect(result.token).toEqual(expect.anything());
            expect(result.instanceServers).toEqual(expect.anything());
            console.log(resp);
        });
    });
});
