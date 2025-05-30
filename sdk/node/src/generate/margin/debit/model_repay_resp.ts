// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, Exclude, plainToClassFromExist } from 'class-transformer';
import { RestResponse } from '@model/common';
import { Response } from '@internal/interfaces/serializable';

export class RepayResp implements Response<RestResponse> {
    /**
     *
     */
    timestamp: number;

    /**
     * Repay Order Id
     */
    orderNo: string;

    /**
     * Actual repay amount
     */
    actualSize: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.timestamp = null;
        // @ts-ignore
        this.orderNo = null;
        // @ts-ignore
        this.actualSize = null;
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
    static fromJson(input: string): RepayResp {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): RepayResp {
        return plainToClassFromExist(new RepayResp(), jsonObject);
    }
}
