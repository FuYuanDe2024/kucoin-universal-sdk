// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { Serializable } from '@internal/interfaces/serializable';

export class GetETHStakingProductsReq implements Serializable {
    /**
     * currency
     */
    currency?: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {}
    /**
     * Creates a new instance of the `GetETHStakingProductsReq` class.
     * The builder pattern allows step-by-step construction of a `GetETHStakingProductsReq` object.
     */
    static builder(): GetETHStakingProductsReqBuilder {
        return new GetETHStakingProductsReqBuilder(new GetETHStakingProductsReq());
    }

    /**
     * Creates a new instance of the `GetETHStakingProductsReq` class with the given data.
     */
    static create(data: {
        /**
         * currency
         */
        currency?: string;
    }): GetETHStakingProductsReq {
        let obj = new GetETHStakingProductsReq();
        obj.currency = data.currency;
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
    static fromJson(input: string): GetETHStakingProductsReq {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetETHStakingProductsReq {
        return plainToClassFromExist(new GetETHStakingProductsReq(), jsonObject);
    }
}

export class GetETHStakingProductsReqBuilder {
    constructor(readonly obj: GetETHStakingProductsReq) {
        this.obj = obj;
    }
    /**
     * currency
     */
    setCurrency(value: string): GetETHStakingProductsReqBuilder {
        this.obj.currency = value;
        return this;
    }

    /**
     * Get the final object.
     */
    build(): GetETHStakingProductsReq {
        return this.obj;
    }
}
