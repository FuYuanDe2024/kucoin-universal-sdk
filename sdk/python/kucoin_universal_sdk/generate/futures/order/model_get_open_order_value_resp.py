# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional
from kucoin_universal_sdk.internal.interfaces.response import Response
from kucoin_universal_sdk.model.common import RestResponse


class GetOpenOrderValueResp(BaseModel, Response):
    """
    GetOpenOrderValueResp

    Attributes:
        open_order_buy_size (int): Total number of the unexecuted buy orders 
        open_order_sell_size (int): Total number of the unexecuted sell orders 
        open_order_buy_cost (str): Value of all the unexecuted buy orders 
        open_order_sell_cost (str): Value of all the unexecuted sell orders 
        settle_currency (str): settlement currency 
    """

    common_response: Optional[RestResponse] = Field(
        default=None, description="Common response")
    open_order_buy_size: Optional[int] = Field(
        default=None,
        description="Total number of the unexecuted buy orders ",
        alias="openOrderBuySize")
    open_order_sell_size: Optional[int] = Field(
        default=None,
        description="Total number of the unexecuted sell orders ",
        alias="openOrderSellSize")
    open_order_buy_cost: Optional[str] = Field(
        default=None,
        description="Value of all the unexecuted buy orders ",
        alias="openOrderBuyCost")
    open_order_sell_cost: Optional[str] = Field(
        default=None,
        description="Value of all the unexecuted sell orders ",
        alias="openOrderSellCost")
    settle_currency: Optional[str] = Field(default=None,
                                           description="settlement currency ",
                                           alias="settleCurrency")

    __properties: ClassVar[List[str]] = [
        "openOrderBuySize", "openOrderSellSize", "openOrderBuyCost",
        "openOrderSellCost", "settleCurrency"
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
    def from_json(cls, json_str: str) -> Optional[GetOpenOrderValueResp]:
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
                                    Any]]) -> Optional[GetOpenOrderValueResp]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "openOrderBuySize":
            obj.get("openOrderBuySize"),
            "openOrderSellSize":
            obj.get("openOrderSellSize"),
            "openOrderBuyCost":
            obj.get("openOrderBuyCost"),
            "openOrderSellCost":
            obj.get("openOrderSellCost"),
            "settleCurrency":
            obj.get("settleCurrency")
        })
        return _obj

    def set_common_response(self, response: RestResponse):
        self.common_response = response
