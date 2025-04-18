// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { GetIsolatedMarginAccountAssetsBaseAsset } from './model_get_isolated_margin_account_assets_base_asset';
import { Type, instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { GetIsolatedMarginAccountAssetsQuoteAsset } from './model_get_isolated_margin_account_assets_quote_asset';
import { Serializable } from '@internal/interfaces/serializable';

export class GetIsolatedMarginAccountAssets implements Serializable {
    /**
     * Symbol
     */
    symbol: string;

    /**
     * Position status; EFFECTIVE-effective, BANKRUPTCY-bankruptcy liquidation, LIQUIDATION-closing, REPAY-repayment, BORROW borrowing
     */
    status: GetIsolatedMarginAccountAssets.StatusEnum;

    /**
     * debt ratio
     */
    debtRatio: string;

    /**
     *
     */
    @Type(() => GetIsolatedMarginAccountAssetsBaseAsset)
    baseAsset: GetIsolatedMarginAccountAssetsBaseAsset;

    /**
     *
     */
    @Type(() => GetIsolatedMarginAccountAssetsQuoteAsset)
    quoteAsset: GetIsolatedMarginAccountAssetsQuoteAsset;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.symbol = null;
        // @ts-ignore
        this.status = null;
        // @ts-ignore
        this.debtRatio = null;
        // @ts-ignore
        this.baseAsset = null;
        // @ts-ignore
        this.quoteAsset = null;
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
    static fromJson(input: string): GetIsolatedMarginAccountAssets {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetIsolatedMarginAccountAssets {
        return plainToClassFromExist(new GetIsolatedMarginAccountAssets(), jsonObject);
    }
}

export namespace GetIsolatedMarginAccountAssets {
    export enum StatusEnum {
        /**
         * Effective
         */
        EFFECTIVE = <any>'EFFECTIVE',
        /**
         * Bankruptcy liquidation
         */
        BANKRUPTCY = <any>'BANKRUPTCY',
        /**
         * Closing
         */
        LIQUIDATION = <any>'LIQUIDATION',
        /**
         * Repayment
         */
        REPAY = <any>'REPAY',
        /**
         * Borrowing
         */
        BORROW = <any>'BORROW',
    }
}
