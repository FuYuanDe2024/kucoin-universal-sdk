// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { Type, instanceToPlain, Exclude, plainToClassFromExist } from 'class-transformer';
import { GetFuturesSubAccountListV2Summary } from './model_get_futures_sub_account_list_v2_summary';
import { GetFuturesSubAccountListV2Accounts } from './model_get_futures_sub_account_list_v2_accounts';
import { RestResponse } from '@model/common';
import { Response } from '@internal/interfaces/serializable';

export class GetFuturesSubAccountListV2Resp implements Response<RestResponse> {
    /**
     *
     */
    @Type(() => GetFuturesSubAccountListV2Summary)
    summary: GetFuturesSubAccountListV2Summary;

    /**
     * Account List
     */
    @Type(() => GetFuturesSubAccountListV2Accounts)
    accounts: Array<GetFuturesSubAccountListV2Accounts>;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.summary = null;
        // @ts-ignore
        this.accounts = null;
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
    static fromJson(input: string): GetFuturesSubAccountListV2Resp {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetFuturesSubAccountListV2Resp {
        return plainToClassFromExist(new GetFuturesSubAccountListV2Resp(), jsonObject);
    }
}
