# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from enum import Enum
from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional
from typing_extensions import Annotated


class GetClosedOrdersReq(BaseModel):
    """
    GetClosedOrdersReq

    Attributes:
        symbol (str): symbol
        trade_type (TradeTypeEnum): Transaction type: MARGIN_TRADE - cross margin trade, MARGIN_ISOLATED_TRADE - isolated margin trade
        side (SideEnum): specify if the order is to 'buy' or 'sell'
        type (TypeEnum): specify if the order is an 'limit' order or 'market' order. 
        last_id (int): The id of the last set of data from the previous batch of data. By default, the latest information is given. lastId is used to filter data and paginate. If lastId is not entered, the default is a maximum of 100 returned data items. The return results include lastId，which can be used as a query parameter to look up new data from the next page.
        limit (int): Default20，Max100
        start_at (int): Start time (milisecond)
        end_at (int): End time (milisecond)
    """

    class TradeTypeEnum(Enum):
        """
        Attributes:
            MARGIN_TRADE: 
            MARGIN_ISOLATED_TRADE: 
        """
        MARGIN_TRADE = 'MARGIN_TRADE'
        MARGIN_ISOLATED_TRADE = 'MARGIN_ISOLATED_TRADE'

    class SideEnum(Enum):
        """
        Attributes:
            BUY: 
            SELL: 
        """
        BUY = 'buy'
        SELL = 'sell'

    class TypeEnum(Enum):
        """
        Attributes:
            LIMIT: 
            MARKET: 
        """
        LIMIT = 'limit'
        MARKET = 'market'

    symbol: Optional[str] = Field(default=None, description="symbol")
    trade_type: Optional[TradeTypeEnum] = Field(
        default=None,
        description=
        "Transaction type: MARGIN_TRADE - cross margin trade, MARGIN_ISOLATED_TRADE - isolated margin trade",
        alias="tradeType")
    side: Optional[SideEnum] = Field(
        default=None, description="specify if the order is to 'buy' or 'sell'")
    type: Optional[TypeEnum] = Field(
        default=None,
        description=
        "specify if the order is an 'limit' order or 'market' order. ")
    last_id: Optional[int] = Field(
        default=None,
        description=
        "The id of the last set of data from the previous batch of data. By default, the latest information is given. lastId is used to filter data and paginate. If lastId is not entered, the default is a maximum of 100 returned data items. The return results include lastId，which can be used as a query parameter to look up new data from the next page.",
        alias="lastId")
    limit: Optional[Annotated[int, Field(le=100, strict=True, ge=1)]] = Field(
        default=20, description="Default20，Max100")
    start_at: Optional[int] = Field(default=None,
                                    description="Start time (milisecond)",
                                    alias="startAt")
    end_at: Optional[int] = Field(default=None,
                                  description="End time (milisecond)",
                                  alias="endAt")

    __properties: ClassVar[List[str]] = [
        "symbol", "tradeType", "side", "type", "lastId", "limit", "startAt",
        "endAt"
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
    def from_json(cls, json_str: str) -> Optional[GetClosedOrdersReq]:
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        _dict = self.model_dump(
            by_alias=True,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(
            cls, obj: Optional[Dict[str,
                                    Any]]) -> Optional[GetClosedOrdersReq]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "symbol":
            obj.get("symbol"),
            "tradeType":
            obj.get("tradeType"),
            "side":
            obj.get("side"),
            "type":
            obj.get("type"),
            "lastId":
            obj.get("lastId"),
            "limit":
            obj.get("limit") if obj.get("limit") is not None else 20,
            "startAt":
            obj.get("startAt"),
            "endAt":
            obj.get("endAt")
        })
        return _obj


class GetClosedOrdersReqBuilder:

    def __init__(self):
        self.obj = {}

    def set_symbol(self, value: str) -> GetClosedOrdersReqBuilder:
        """
        symbol
        """
        self.obj['symbol'] = value
        return self

    def set_trade_type(
            self, value: GetClosedOrdersReq.TradeTypeEnum
    ) -> GetClosedOrdersReqBuilder:
        """
        Transaction type: MARGIN_TRADE - cross margin trade, MARGIN_ISOLATED_TRADE - isolated margin trade
        """
        self.obj['tradeType'] = value
        return self

    def set_side(
            self,
            value: GetClosedOrdersReq.SideEnum) -> GetClosedOrdersReqBuilder:
        """
        specify if the order is to 'buy' or 'sell'
        """
        self.obj['side'] = value
        return self

    def set_type(
            self,
            value: GetClosedOrdersReq.TypeEnum) -> GetClosedOrdersReqBuilder:
        """
        specify if the order is an 'limit' order or 'market' order. 
        """
        self.obj['type'] = value
        return self

    def set_last_id(self, value: int) -> GetClosedOrdersReqBuilder:
        """
        The id of the last set of data from the previous batch of data. By default, the latest information is given. lastId is used to filter data and paginate. If lastId is not entered, the default is a maximum of 100 returned data items. The return results include lastId，which can be used as a query parameter to look up new data from the next page.
        """
        self.obj['lastId'] = value
        return self

    def set_limit(self, value: int) -> GetClosedOrdersReqBuilder:
        """
        Default20，Max100
        """
        self.obj['limit'] = value
        return self

    def set_start_at(self, value: int) -> GetClosedOrdersReqBuilder:
        """
        Start time (milisecond)
        """
        self.obj['startAt'] = value
        return self

    def set_end_at(self, value: int) -> GetClosedOrdersReqBuilder:
        """
        End time (milisecond)
        """
        self.obj['endAt'] = value
        return self

    def build(self) -> GetClosedOrdersReq:
        return GetClosedOrdersReq(**self.obj)
