# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from enum import Enum
from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional


class AddOcoOrderReq(BaseModel):
    """
    AddOcoOrderReq

    Attributes:
        client_oid (str): Client Order Id，The ClientOid field is a unique ID created by the user（we recommend using a UUID）, and can only contain numbers, letters, underscores （_）, and hyphens （-）. This field is returned when order information is obtained. You can use clientOid to tag your orders. ClientOid is different from the order ID created by the service provider. Please do not initiate requests using the same clientOid. The maximum length for the ClientOid is 40 characters.  Please remember the orderId created by the service provider, it used to check for updates in order status.
        side (SideEnum): specify if the order is to 'buy' or 'sell'
        symbol (str): symbol
        remark (str): Order placement remarks, length cannot exceed 20 characters (ASCII)
        price (str): Specify price for order
        size (str): Specify quantity for order
        stop_price (str): trigger price.
        limit_price (str): The limit order price after take-profit and stop-loss are triggered.
        trade_type (TradeTypeEnum): Transaction Type, currently only supports TRADE (spot transactions), the default is TRADE
    """

    class SideEnum(Enum):
        """
        Attributes:
            BUY: 
            SELL: 
        """
        BUY = 'buy'
        SELL = 'sell'

    class TradeTypeEnum(Enum):
        """
        Attributes:
            TRADE: Spot Trading
        """
        TRADE = 'TRADE'

    client_oid: Optional[str] = Field(
        default=None,
        description=
        "Client Order Id，The ClientOid field is a unique ID created by the user（we recommend using a UUID）, and can only contain numbers, letters, underscores （_）, and hyphens （-）. This field is returned when order information is obtained. You can use clientOid to tag your orders. ClientOid is different from the order ID created by the service provider. Please do not initiate requests using the same clientOid. The maximum length for the ClientOid is 40 characters.  Please remember the orderId created by the service provider, it used to check for updates in order status.",
        alias="clientOid")
    side: Optional[SideEnum] = Field(
        default=None, description="specify if the order is to 'buy' or 'sell'")
    symbol: Optional[str] = Field(default=None, description="symbol")
    remark: Optional[str] = Field(
        default=None,
        description=
        "Order placement remarks, length cannot exceed 20 characters (ASCII)")
    price: Optional[str] = Field(default=None,
                                 description="Specify price for order")
    size: Optional[str] = Field(default=None,
                                description="Specify quantity for order")
    stop_price: Optional[str] = Field(default=None,
                                      description="trigger price.",
                                      alias="stopPrice")
    limit_price: Optional[str] = Field(
        default=None,
        description=
        "The limit order price after take-profit and stop-loss are triggered.",
        alias="limitPrice")
    trade_type: Optional[TradeTypeEnum] = Field(
        default=None,
        description=
        "Transaction Type, currently only supports TRADE (spot transactions), the default is TRADE",
        alias="tradeType")

    __properties: ClassVar[List[str]] = [
        "clientOid", "side", "symbol", "remark", "price", "size", "stopPrice",
        "limitPrice", "tradeType"
    ]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=False,
        protected_namespaces=(),
    )

    def to_str(self) -> str:
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_none=True)

    @classmethod
    def from_json(cls, json_str: str) -> Optional[AddOcoOrderReq]:
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        _dict = self.model_dump(
            by_alias=True,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str,
                                          Any]]) -> Optional[AddOcoOrderReq]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "clientOid": obj.get("clientOid"),
            "side": obj.get("side"),
            "symbol": obj.get("symbol"),
            "remark": obj.get("remark"),
            "price": obj.get("price"),
            "size": obj.get("size"),
            "stopPrice": obj.get("stopPrice"),
            "limitPrice": obj.get("limitPrice"),
            "tradeType": obj.get("tradeType")
        })
        return _obj


class AddOcoOrderReqBuilder:

    def __init__(self):
        self.obj = {}

    def set_client_oid(self, value: str) -> AddOcoOrderReqBuilder:
        """
        Client Order Id，The ClientOid field is a unique ID created by the user（we recommend using a UUID）, and can only contain numbers, letters, underscores （_）, and hyphens （-）. This field is returned when order information is obtained. You can use clientOid to tag your orders. ClientOid is different from the order ID created by the service provider. Please do not initiate requests using the same clientOid. The maximum length for the ClientOid is 40 characters.  Please remember the orderId created by the service provider, it used to check for updates in order status.
        """
        self.obj['clientOid'] = value
        return self

    def set_side(self,
                 value: AddOcoOrderReq.SideEnum) -> AddOcoOrderReqBuilder:
        """
        specify if the order is to 'buy' or 'sell'
        """
        self.obj['side'] = value
        return self

    def set_symbol(self, value: str) -> AddOcoOrderReqBuilder:
        """
        symbol
        """
        self.obj['symbol'] = value
        return self

    def set_remark(self, value: str) -> AddOcoOrderReqBuilder:
        """
        Order placement remarks, length cannot exceed 20 characters (ASCII)
        """
        self.obj['remark'] = value
        return self

    def set_price(self, value: str) -> AddOcoOrderReqBuilder:
        """
        Specify price for order
        """
        self.obj['price'] = value
        return self

    def set_size(self, value: str) -> AddOcoOrderReqBuilder:
        """
        Specify quantity for order
        """
        self.obj['size'] = value
        return self

    def set_stop_price(self, value: str) -> AddOcoOrderReqBuilder:
        """
        trigger price.
        """
        self.obj['stopPrice'] = value
        return self

    def set_limit_price(self, value: str) -> AddOcoOrderReqBuilder:
        """
        The limit order price after take-profit and stop-loss are triggered.
        """
        self.obj['limitPrice'] = value
        return self

    def set_trade_type(
            self,
            value: AddOcoOrderReq.TradeTypeEnum) -> AddOcoOrderReqBuilder:
        """
        Transaction Type, currently only supports TRADE (spot transactions), the default is TRADE
        """
        self.obj['tradeType'] = value
        return self

    def build(self) -> AddOcoOrderReq:
        return AddOcoOrderReq(**self.obj)
