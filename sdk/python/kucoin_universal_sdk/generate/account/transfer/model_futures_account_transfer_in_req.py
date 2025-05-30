# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from enum import Enum
from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional


class FuturesAccountTransferInReq(BaseModel):
    """
    FuturesAccountTransferInReq

    Attributes:
        currency (str): Currency, including XBT,USDT...
        amount (float): Amount to be transfered in
        pay_account_type (PayAccountTypeEnum): Payment account type, including MAIN,TRADE
    """

    class PayAccountTypeEnum(Enum):
        """
        Attributes:
            MAIN: 
            TRADE: 
        """
        MAIN = 'MAIN'
        TRADE = 'TRADE'

    currency: Optional[str] = Field(
        default=None, description="Currency, including XBT,USDT...")
    amount: Optional[float] = Field(default=None,
                                    description="Amount to be transfered in")
    pay_account_type: Optional[PayAccountTypeEnum] = Field(
        default=None,
        description="Payment account type, including MAIN,TRADE",
        alias="payAccountType")

    __properties: ClassVar[List[str]] = [
        "currency", "amount", "payAccountType"
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
    def from_json(cls, json_str: str) -> Optional[FuturesAccountTransferInReq]:
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        _dict = self.model_dump(
            by_alias=True,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(
        cls,
        obj: Optional[Dict[str,
                           Any]]) -> Optional[FuturesAccountTransferInReq]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "currency": obj.get("currency"),
            "amount": obj.get("amount"),
            "payAccountType": obj.get("payAccountType")
        })
        return _obj


class FuturesAccountTransferInReqBuilder:

    def __init__(self):
        self.obj = {}

    def set_currency(self, value: str) -> FuturesAccountTransferInReqBuilder:
        """
        Currency, including XBT,USDT...
        """
        self.obj['currency'] = value
        return self

    def set_amount(self, value: float) -> FuturesAccountTransferInReqBuilder:
        """
        Amount to be transfered in
        """
        self.obj['amount'] = value
        return self

    def set_pay_account_type(
        self, value: FuturesAccountTransferInReq.PayAccountTypeEnum
    ) -> FuturesAccountTransferInReqBuilder:
        """
        Payment account type, including MAIN,TRADE
        """
        self.obj['payAccountType'] = value
        return self

    def build(self) -> FuturesAccountTransferInReq:
        return FuturesAccountTransferInReq(**self.obj)
