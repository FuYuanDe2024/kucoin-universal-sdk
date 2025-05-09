// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { GetBorrowHistoryItems } from './model_get_borrow_history_items';
import { Type, instanceToPlain, Exclude, plainToClassFromExist } from 'class-transformer';
import { RestResponse } from '@model/common';
import { Response } from '@internal/interfaces/serializable';

export class GetBorrowHistoryResp implements Response<RestResponse> {
    /**
     *
     */
    timestamp: number;

    /**
     * current page
     */
    currentPage: number;

    /**
     * page size
     */
    pageSize: number;

    /**
     * total number
     */
    totalNum: number;

    /**
     * total page
     */
    totalPage: number;

    /**
     *
     */
    @Type(() => GetBorrowHistoryItems)
    items: Array<GetBorrowHistoryItems>;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.timestamp = null;
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
    static fromJson(input: string): GetBorrowHistoryResp {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetBorrowHistoryResp {
        return plainToClassFromExist(new GetBorrowHistoryResp(), jsonObject);
    }
}
