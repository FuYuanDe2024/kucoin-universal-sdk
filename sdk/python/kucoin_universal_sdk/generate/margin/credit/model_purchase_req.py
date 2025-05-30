# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional


class PurchaseReq(BaseModel):
    """
    PurchaseReq

    Attributes:
        currency (str): Currency
        size (str): purchase amount
        interest_rate (str): purchase interest rate
    """

    currency: Optional[str] = Field(default=None, description="Currency")
    size: Optional[str] = Field(default=None, description="purchase amount")
    interest_rate: Optional[str] = Field(default=None,
                                         description="purchase interest rate",
                                         alias="interestRate")

    __properties: ClassVar[List[str]] = ["currency", "size", "interestRate"]

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
    def from_json(cls, json_str: str) -> Optional[PurchaseReq]:
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        _dict = self.model_dump(
            by_alias=True,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[PurchaseReq]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "currency": obj.get("currency"),
            "size": obj.get("size"),
            "interestRate": obj.get("interestRate")
        })
        return _obj


class PurchaseReqBuilder:

    def __init__(self):
        self.obj = {}

    def set_currency(self, value: str) -> PurchaseReqBuilder:
        """
        Currency
        """
        self.obj['currency'] = value
        return self

    def set_size(self, value: str) -> PurchaseReqBuilder:
        """
        purchase amount
        """
        self.obj['size'] = value
        return self

    def set_interest_rate(self, value: str) -> PurchaseReqBuilder:
        """
        purchase interest rate
        """
        self.obj['interestRate'] = value
        return self

    def build(self) -> PurchaseReq:
        return PurchaseReq(**self.obj)
