// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { Serializable } from '@internal/interfaces/serializable';

export class GetRedeemOrdersReq implements Serializable {
    /**
     * currency
     */
    currency?: string;

    /**
     * DONE-completed; PENDING-settling
     */
    status?: GetRedeemOrdersReq.StatusEnum;

    /**
     * Redeem order id
     */
    redeemOrderNo?: string;

    /**
     * Current page; default is 1
     */
    currentPage?: number = 1;

    /**
     * Page size; 1<=pageSize<=100; default is 50
     */
    pageSize?: number = 50;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {}
    /**
     * Creates a new instance of the `GetRedeemOrdersReq` class.
     * The builder pattern allows step-by-step construction of a `GetRedeemOrdersReq` object.
     */
    static builder(): GetRedeemOrdersReqBuilder {
        return new GetRedeemOrdersReqBuilder(new GetRedeemOrdersReq());
    }

    /**
     * Creates a new instance of the `GetRedeemOrdersReq` class with the given data.
     */
    static create(data: {
        /**
         * currency
         */
        currency?: string;
        /**
         * DONE-completed; PENDING-settling
         */
        status?: GetRedeemOrdersReq.StatusEnum;
        /**
         * Redeem order id
         */
        redeemOrderNo?: string;
        /**
         * Current page; default is 1
         */
        currentPage?: number;
        /**
         * Page size; 1<=pageSize<=100; default is 50
         */
        pageSize?: number;
    }): GetRedeemOrdersReq {
        let obj = new GetRedeemOrdersReq();
        obj.currency = data.currency;
        obj.status = data.status;
        obj.redeemOrderNo = data.redeemOrderNo;
        if (data.currentPage) {
            obj.currentPage = data.currentPage;
        } else {
            obj.currentPage = 1;
        }
        if (data.pageSize) {
            obj.pageSize = data.pageSize;
        } else {
            obj.pageSize = 50;
        }
        return obj;
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
    static fromJson(input: string): GetRedeemOrdersReq {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetRedeemOrdersReq {
        return plainToClassFromExist(new GetRedeemOrdersReq(), jsonObject);
    }
}

export namespace GetRedeemOrdersReq {
    export enum StatusEnum {
        /**
         *
         */
        DONE = <any>'DONE',
        /**
         *
         */
        PENDING = <any>'PENDING',
    }
}

export class GetRedeemOrdersReqBuilder {
    constructor(readonly obj: GetRedeemOrdersReq) {
        this.obj = obj;
    }
    /**
     * currency
     */
    setCurrency(value: string): GetRedeemOrdersReqBuilder {
        this.obj.currency = value;
        return this;
    }

    /**
     * DONE-completed; PENDING-settling
     */
    setStatus(value: GetRedeemOrdersReq.StatusEnum): GetRedeemOrdersReqBuilder {
        this.obj.status = value;
        return this;
    }

    /**
     * Redeem order id
     */
    setRedeemOrderNo(value: string): GetRedeemOrdersReqBuilder {
        this.obj.redeemOrderNo = value;
        return this;
    }

    /**
     * Current page; default is 1
     */
    setCurrentPage(value: number): GetRedeemOrdersReqBuilder {
        this.obj.currentPage = value;
        return this;
    }

    /**
     * Page size; 1<=pageSize<=100; default is 50
     */
    setPageSize(value: number): GetRedeemOrdersReqBuilder {
        this.obj.pageSize = value;
        return this;
    }

    /**
     * Get the final object.
     */
    build(): GetRedeemOrdersReq {
        return this.obj;
    }
}
