// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { Type, instanceToPlain, Exclude, plainToClassFromExist } from 'class-transformer';
import { GetTradeHistoryOldItems } from './model_get_trade_history_old_items';
import { RestResponse } from '@model/common';
import { Response } from '@internal/interfaces/serializable';

export class GetTradeHistoryOldResp implements Response<RestResponse> {
    /**
     *
     */
    currentPage: number;

    /**
     *
     */
    pageSize: number;

    /**
     *
     */
    totalNum: number;

    /**
     *
     */
    totalPage: number;

    /**
     *
     */
    @Type(() => GetTradeHistoryOldItems)
    items: Array<GetTradeHistoryOldItems>;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.currentPage = null;
        // @ts-ignore
        this.pageSize = null;
        // @ts-ignore
        this.totalNum = null;
        // @ts-ignore
        this.totalPage = null;
        // @ts-ignore
        this.items = null;
    }
    /**
     * common response
     */
    @Exclude()
    commonResponse?: RestResponse;

    setCommonResponse(response: RestResponse): void {
        this.commonResponse = response;
    }

    /**
     * Convert the object to a JSON string.
     */
    toJson(): string {
        return JSON.stringify(instanceToPlain(this));
    }
    /**
     * Create an object from a JSON string.
     */
    static fromJson(input: string): GetTradeHistoryOldResp {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetTradeHistoryOldResp {
        return plainToClassFromExist(new GetTradeHistoryOldResp(), jsonObject);
    }
}
