// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { Serializable } from '@internal/interfaces/serializable';

export class GetDepositAddressV3Req implements Serializable {
    /**
     * currency
     */
    currency?: string;

    /**
     * Deposit amount. This parameter is only used when applying for invoices on the Lightning Network. This parameter is invalid if it is not passed through the Lightning Network.
     */
    amount?: string;

    /**
     * The chain Id of currency.
     */
    chain?: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {}
    /**
     * Creates a new instance of the `GetDepositAddressV3Req` class.
     * The builder pattern allows step-by-step construction of a `GetDepositAddressV3Req` object.
     */
    static builder(): GetDepositAddressV3ReqBuilder {
        return new GetDepositAddressV3ReqBuilder(new GetDepositAddressV3Req());
    }

    /**
     * Creates a new instance of the `GetDepositAddressV3Req` class with the given data.
     */
    static create(data: {
        /**
         * currency
         */
        currency?: string;
        /**
         * Deposit amount. This parameter is only used when applying for invoices on the Lightning Network. This parameter is invalid if it is not passed through the Lightning Network.
         */
        amount?: string;
        /**
         * The chain Id of currency.
         */
        chain?: string;
    }): GetDepositAddressV3Req {
        let obj = new GetDepositAddressV3Req();
        obj.currency = data.currency;
        obj.amount = data.amount;
        obj.chain = data.chain;
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
    static fromJson(input: string): GetDepositAddressV3Req {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): GetDepositAddressV3Req {
        return plainToClassFromExist(new GetDepositAddressV3Req(), jsonObject);
    }
}

export class GetDepositAddressV3ReqBuilder {
    constructor(readonly obj: GetDepositAddressV3Req) {
        this.obj = obj;
    }
    /**
     * currency
     */
    setCurrency(value: string): GetDepositAddressV3ReqBuilder {
        this.obj.currency = value;
        return this;
    }

    /**
     * Deposit amount. This parameter is only used when applying for invoices on the Lightning Network. This parameter is invalid if it is not passed through the Lightning Network.
     */
    setAmount(value: string): GetDepositAddressV3ReqBuilder {
        this.obj.amount = value;
        return this;
    }

    /**
     * The chain Id of currency.
     */
    setChain(value: string): GetDepositAddressV3ReqBuilder {
        this.obj.chain = value;
        return this;
    }

    /**
     * Get the final object.
     */
    build(): GetDepositAddressV3Req {
        return this.obj;
    }
}
