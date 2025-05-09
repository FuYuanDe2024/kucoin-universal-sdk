// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { Serializable } from '@internal/interfaces/serializable';

export class ModifyOrderReq implements Serializable {
    /**
     * The old client order id,orderId and clientOid must choose one
     */
    clientOid?: string;

    /**
     * symbol
     */
    symbol: string;

    /**
     * The old order id, orderId and clientOid must choose one
     */
    orderId?: string;

    /**
     * The modified price of the new order, newPrice and newSize must choose one
     */
    newPrice?: string;

    /**
     * The modified size of the new order, newPrice and newSize must choose one
     */
    newSize?: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.symbol = null;
    }
    /**
     * Creates a new instance of the `ModifyOrderReq` class.
     * The builder pattern allows step-by-step construction of a `ModifyOrderReq` object.
     */
    static builder(): ModifyOrderReqBuilder {
        return new ModifyOrderReqBuilder(new ModifyOrderReq());
    }

    /**
     * Creates a new instance of the `ModifyOrderReq` class with the given data.
     */
    static create(data: {
        /**
         * The old client order id,orderId and clientOid must choose one
         */
        clientOid?: string;
        /**
         * symbol
         */
        symbol: string;
        /**
         * The old order id, orderId and clientOid must choose one
         */
        orderId?: string;
        /**
         * The modified price of the new order, newPrice and newSize must choose one
         */
        newPrice?: string;
        /**
         * The modified size of the new order, newPrice and newSize must choose one
         */
        newSize?: string;
    }): ModifyOrderReq {
        let obj = new ModifyOrderReq();
        obj.clientOid = data.clientOid;
        obj.symbol = data.symbol;
        obj.orderId = data.orderId;
        obj.newPrice = data.newPrice;
        obj.newSize = data.newSize;
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
    static fromJson(input: string): ModifyOrderReq {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): ModifyOrderReq {
        return plainToClassFromExist(new ModifyOrderReq(), jsonObject);
    }
}

export class ModifyOrderReqBuilder {
    constructor(readonly obj: ModifyOrderReq) {
        this.obj = obj;
    }
    /**
     * The old client order id,orderId and clientOid must choose one
     */
    setClientOid(value: string): ModifyOrderReqBuilder {
        this.obj.clientOid = value;
        return this;
    }

    /**
     * symbol
     */
    setSymbol(value: string): ModifyOrderReqBuilder {
        this.obj.symbol = value;
        return this;
    }

    /**
     * The old order id, orderId and clientOid must choose one
     */
    setOrderId(value: string): ModifyOrderReqBuilder {
        this.obj.orderId = value;
        return this;
    }

    /**
     * The modified price of the new order, newPrice and newSize must choose one
     */
    setNewPrice(value: string): ModifyOrderReqBuilder {
        this.obj.newPrice = value;
        return this;
    }

    /**
     * The modified size of the new order, newPrice and newSize must choose one
     */
    setNewSize(value: string): ModifyOrderReqBuilder {
        this.obj.newSize = value;
        return this;
    }

    /**
     * Get the final object.
     */
    build(): ModifyOrderReq {
        return this.obj;
    }
}
