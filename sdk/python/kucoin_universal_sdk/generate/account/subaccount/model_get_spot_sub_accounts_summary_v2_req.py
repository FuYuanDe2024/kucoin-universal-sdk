# coding: utf-8

# Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

from __future__ import annotations
import pprint
import json

from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional
from typing_extensions import Annotated


class GetSpotSubAccountsSummaryV2Req(BaseModel):
    """
    GetSpotSubAccountsSummaryV2Req

    Attributes:
        current_page (int): Current request page. Default is 1
        page_size (int): Number of results per request. Minimum is 1, maximum is 100, default is 10.
    """

    current_page: Optional[int] = Field(
        default=None,
        description="Current request page. Default is 1",
        alias="currentPage")
    page_size: Optional[Annotated[int, Field(
        le=100, strict=True, ge=1
    )]] = Field(
        default=10,
        description=
        "Number of results per request. Minimum is 1, maximum is 100, default is 10.",
        alias="pageSize")

    __properties: ClassVar[List[str]] = ["currentPage", "pageSize"]

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
    def from_json(cls,
                  json_str: str) -> Optional[GetSpotSubAccountsSummaryV2Req]:
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
                           Any]]) -> Optional[GetSpotSubAccountsSummaryV2Req]:
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "currentPage":
            obj.get("currentPage"),
            "pageSize":
            obj.get("pageSize") if obj.get("pageSize") is not None else 10
        })
        return _obj


class GetSpotSubAccountsSummaryV2ReqBuilder:

    def __init__(self):
        self.obj = {}

    def set_current_page(self,
                         value: int) -> GetSpotSubAccountsSummaryV2ReqBuilder:
        """
        Current request page. Default is 1
        """
        self.obj['currentPage'] = value
        return self

    def set_page_size(self,
                      value: int) -> GetSpotSubAccountsSummaryV2ReqBuilder:
        """
        Number of results per request. Minimum is 1, maximum is 100, default is 10.
        """
        self.obj['pageSize'] = value
        return self

    def build(self) -> GetSpotSubAccountsSummaryV2Req:
        return GetSpotSubAccountsSummaryV2Req(**self.obj)
