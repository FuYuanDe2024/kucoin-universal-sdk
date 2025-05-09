# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from abc import ABC, abstractmethod
from kucoin_universal_sdk.internal.interfaces.websocket import WebSocketService
from .model_all_tickers_event import AllTickersEventCallback, AllTickersEventCallbackWrapper
from .model_klines_event import KlinesEventCallback, KlinesEventCallbackWrapper
from .model_market_snapshot_event import MarketSnapshotEventCallback, MarketSnapshotEventCallbackWrapper
from .model_orderbook_increment_event import OrderbookIncrementEventCallback, OrderbookIncrementEventCallbackWrapper
from .model_orderbook_level1_event import OrderbookLevel1EventCallback, OrderbookLevel1EventCallbackWrapper
from .model_orderbook_level50_event import OrderbookLevel50EventCallback, OrderbookLevel50EventCallbackWrapper
from .model_orderbook_level5_event import OrderbookLevel5EventCallback, OrderbookLevel5EventCallbackWrapper
from .model_symbol_snapshot_event import SymbolSnapshotEventCallback, SymbolSnapshotEventCallbackWrapper
from .model_ticker_event import TickerEventCallback, TickerEventCallbackWrapper
from .model_trade_event import TradeEventCallback, TradeEventCallbackWrapper
from typing import List


class SpotPublicWS(ABC):

    @abstractmethod
    def all_tickers(self, callback: AllTickersEventCallback) -> str:
        """
        summary: Get All Tickers
        description: Subscribe to this topic to get the push of all market symbols BBO change.
        push frequency: once every 100ms
        """
        pass

    @abstractmethod
    def klines(self, symbol: str, type: str,
               callback: KlinesEventCallback) -> str:
        """
        summary: Klines
        description: Subscribe to this topic to get K-Line data.
        push frequency: real-time
        """
        pass

    @abstractmethod
    def market_snapshot(self, market: str,
                        callback: MarketSnapshotEventCallback) -> str:
        """
        summary: Market Snapshot
        description: Subscribe this topic to get the snapshot data of for the entire market.
        push frequency: once every 2s
        """
        pass

    @abstractmethod
    def orderbook_increment(self, symbol: List[str],
                            callback: OrderbookIncrementEventCallback) -> str:
        """
        summary: Orderbook - Increment
        description: The system will return the increment change orderbook data(All depth), A topic supports up to 100 symbols. If there is no change in the market, data will not be pushed
        push frequency: real-time
        """
        pass

    @abstractmethod
    def orderbook_level1(self, symbol: List[str],
                         callback: OrderbookLevel1EventCallback) -> str:
        """
        summary: Orderbook - Level1
        description: The system will return the 1 best ask/bid orders data, A topic supports up to 100 symbols. If there is no change in the market, data will not be pushed
        push frequency: once every 10ms
        """
        pass

    @abstractmethod
    def orderbook_level50(self, symbol: List[str],
                          callback: OrderbookLevel50EventCallback) -> str:
        """
        summary: Orderbook - Level50
        description: The system will return the 50 best ask/bid orders data, A topic supports up to 100 symbols. If there is no change in the market, data will not be pushed
        push frequency: once every 100ms
        """
        pass

    @abstractmethod
    def orderbook_level5(self, symbol: List[str],
                         callback: OrderbookLevel5EventCallback) -> str:
        """
        summary: Orderbook - Level5
        description: The system will return the 5 best ask/bid orders data,A topic supports up to 100 symbols. If there is no change in the market, data will not be pushed
        push frequency: once every 100ms
        """
        pass

    @abstractmethod
    def symbol_snapshot(self, symbol: str,
                        callback: SymbolSnapshotEventCallback) -> str:
        """
        summary: Symbol Snapshot
        description: Subscribe to get snapshot data for a single symbol.
        push frequency: once every 2s
        """
        pass

    @abstractmethod
    def ticker(self, symbol: List[str], callback: TickerEventCallback) -> str:
        """
        summary: Get Ticker
        description: Subscribe to this topic to get the specified symbol push of BBO changes.
        push frequency: once every 100ms
        """
        pass

    @abstractmethod
    def trade(self, symbol: List[str], callback: TradeEventCallback) -> str:
        """
        summary: Trade
        description: Subscribe to this topic to get the matching event data flow of Level 3. A topic supports up to 100 symbols.
        push frequency: real-time
        """
        pass

    @abstractmethod
    def unsubscribe(self, id: str):
        pass

    @abstractmethod
    def start(self):
        pass

    @abstractmethod
    def stop(self):
        pass


class SpotPublicWSImpl(SpotPublicWS):

    def __init__(self, transport: WebSocketService):
        self.transport = transport

    def all_tickers(self, callback: AllTickersEventCallback) -> str:
        topic_prefix = "/market/ticker:all"

        args = []

        return self.transport.subscribe(
            topic_prefix, args, AllTickersEventCallbackWrapper(callback))

    def klines(self, symbol: str, type: str,
               callback: KlinesEventCallback) -> str:
        topic_prefix = "/market/candles"

        args = ["_".join([
            symbol,
            type,
        ])]

        return self.transport.subscribe(topic_prefix, args,
                                        KlinesEventCallbackWrapper(callback))

    def market_snapshot(self, market: str,
                        callback: MarketSnapshotEventCallback) -> str:
        topic_prefix = "/market/snapshot"

        args = [market]

        return self.transport.subscribe(
            topic_prefix, args, MarketSnapshotEventCallbackWrapper(callback))

    def orderbook_increment(self, symbol: List[str],
                            callback: OrderbookIncrementEventCallback) -> str:
        topic_prefix = "/market/level2"

        args = symbol

        return self.transport.subscribe(
            topic_prefix, args,
            OrderbookIncrementEventCallbackWrapper(callback))

    def orderbook_level1(self, symbol: List[str],
                         callback: OrderbookLevel1EventCallback) -> str:
        topic_prefix = "/spotMarket/level1"

        args = symbol

        return self.transport.subscribe(
            topic_prefix, args, OrderbookLevel1EventCallbackWrapper(callback))

    def orderbook_level50(self, symbol: List[str],
                          callback: OrderbookLevel50EventCallback) -> str:
        topic_prefix = "/market/level2"

        args = symbol

        return self.transport.subscribe(
            topic_prefix, args, OrderbookLevel50EventCallbackWrapper(callback))

    def orderbook_level5(self, symbol: List[str],
                         callback: OrderbookLevel5EventCallback) -> str:
        topic_prefix = "/spotMarket/level2Depth5"

        args = symbol

        return self.transport.subscribe(
            topic_prefix, args, OrderbookLevel5EventCallbackWrapper(callback))

    def symbol_snapshot(self, symbol: str,
                        callback: SymbolSnapshotEventCallback) -> str:
        topic_prefix = "/market/snapshot"

        args = [symbol]

        return self.transport.subscribe(
            topic_prefix, args, SymbolSnapshotEventCallbackWrapper(callback))

    def ticker(self, symbol: List[str], callback: TickerEventCallback) -> str:
        topic_prefix = "/market/ticker"

        args = symbol

        return self.transport.subscribe(topic_prefix, args,
                                        TickerEventCallbackWrapper(callback))

    def trade(self, symbol: List[str], callback: TradeEventCallback) -> str:
        topic_prefix = "/market/match"

        args = symbol

        return self.transport.subscribe(topic_prefix, args,
                                        TradeEventCallbackWrapper(callback))

    def unsubscribe(self, id: str):
        self.transport.unsubscribe(id)

    def start(self):
        self.transport.start()

    def stop(self):
        self.transport.stop()
