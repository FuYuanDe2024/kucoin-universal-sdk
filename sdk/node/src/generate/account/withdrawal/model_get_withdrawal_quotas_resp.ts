// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, Exclude, plainToClassFromExist } from 'class-transformer';
import { RestResponse } from '@model/common';
import { Response } from '@internal/interfaces/serializable';

export class GetWithdrawalQuotasResp implements Response<RestResponse> {
    /**
     *
     */
    currency: string;

    /**
     *
     */
    limitBTCAmount: string;

    /**
     *
     */
    usedBTCAmount: string;

    /**
     * withdrawal limit currency
     */
    quotaCurrency: string;

    /**
     * The intraday available withdrawal amount(withdrawal limit currency)
     */
    limitQuotaCurrencyAmount: string;

    /**
     * The intraday cumulative withdrawal amount(withdrawal limit currency)
     */
    usedQuotaCurrencyAmount: string;

    /**
     * Remaining amount available to withdraw the current day
     */
    remainAmount: string;

    /**
     * Current available withdrawal amount
     */
    availableAmount: string;

    /**
     * Minimum withdrawal fee
     */
    withdrawMinFee: string;

    /**
     * Fees for internal withdrawal
     */
    innerWithdrawMinFee: string;

    /**
     * Minimum withdrawal amount
     */
    withdrawMinSize: string;

    /**
     * Is the withdraw function enabled or not
     */
    isWithdrawEnabled: boolean;

    /**
     * Floating point precision.
     */
    precision: number;

    /**
     * The chainName of currency
     */
    chain: string;

    /**
     * Reasons for restriction, Usually empty
     */
    reason: string;

    /**
     * Total locked amount (including the amount locked into USDT for each currency)
     */
    lockedAmount: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.currency = null;
        // @ts-ignore
        this.limitBTCAmount = null;
        // @ts-ignore
        this.usedBTCAmount = null;
        // @ts-ignore
        this.quotaCurrency = null;
        // @ts-ignore
        this.limitQuotaCurrencyAmount = null;
        // @ts-ignore
        this.usedQuotaCurrencyAmount = null;
        // @ts-ignore
        this.remainAmount = null;
        // @ts-ignore
        this.availableAmount = null;
        // @ts-ignore
        this.withdrawMinFee = null;
        // @ts-ignore
        this.innerWithdrawMinFee = null;
        // @ts-ignore
        this.withdrawMinSize = null;
        // @ts-ignore
        this.isWithdrawEnabled = null;
        // @ts-ignore
        this.precision = null;
        // @ts-ignore
        this.chain = null;
        // @ts-ignore
        this.reason = null;
        // @ts-ignore
        this.lockedAmount = null;
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
    static fromJson(input: string): GetWithdrawalQuotasResp {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetWithdrawalQuotasResp {
        return plainToClassFromExist(new GetWithdrawalQuotasResp(), jsonObject);
    }
}
