// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { instanceToPlain, plainToClassFromExist } from 'class-transformer';
import { Serializable } from '@internal/interfaces/serializable';

export class AddSubAccountApiReq implements Serializable {
    /**
     * Password(Must contain 7-32 characters. Cannot contain any spaces.)
     */
    passphrase: string;

    /**
     * Remarks(1~24 characters)
     */
    remark: string;

    /**
     * [Permissions](https://www.kucoin.com/docs-new/doc-338144)(Only General、Spot、Futures、Margin、InnerTransfer(Flex Transfer) permissions can be set, such as \"General, Trade\". The default is \"General\")
     */
    permission?: string = 'General';

    /**
     * IP whitelist(You may add up to 20 IPs. Use a halfwidth comma to each IP)
     */
    ipWhitelist?: string;

    /**
     * API expiration time; Never expire(default)-1，30Day30，90Day90，180Day180，360Day360
     */
    expire?: AddSubAccountApiReq.ExpireEnum = AddSubAccountApiReq.ExpireEnum._1;

    /**
     * Sub-account name, create sub account name of API Key.
     */
    subName: string;

    /**
     * Private constructor, please use the corresponding static methods to construct the object.
     */
    private constructor() {
        // @ts-ignore
        this.passphrase = null;
        // @ts-ignore
        this.remark = null;
        // @ts-ignore
        this.subName = null;
    }
    /**
     * Creates a new instance of the `AddSubAccountApiReq` class.
     * The builder pattern allows step-by-step construction of a `AddSubAccountApiReq` object.
     */
    static builder(): AddSubAccountApiReqBuilder {
        return new AddSubAccountApiReqBuilder(new AddSubAccountApiReq());
    }

    /**
     * Creates a new instance of the `AddSubAccountApiReq` class with the given data.
     */
    static create(data: {
        /**
         * Password(Must contain 7-32 characters. Cannot contain any spaces.)
         */
        passphrase: string;
        /**
         * Remarks(1~24 characters)
         */
        remark: string;
        /**
         * [Permissions](https://www.kucoin.com/docs-new/doc-338144)(Only General、Spot、Futures、Margin、InnerTransfer(Flex Transfer) permissions can be set, such as \"General, Trade\". The default is \"General\")
         */
        permission?: string;
        /**
         * IP whitelist(You may add up to 20 IPs. Use a halfwidth comma to each IP)
         */
        ipWhitelist?: string;
        /**
         * API expiration time; Never expire(default)-1，30Day30，90Day90，180Day180，360Day360
         */
        expire?: AddSubAccountApiReq.ExpireEnum;
        /**
         * Sub-account name, create sub account name of API Key.
         */
        subName: string;
    }): AddSubAccountApiReq {
        let obj = new AddSubAccountApiReq();
        obj.passphrase = data.passphrase;
        obj.remark = data.remark;
        if (data.permission) {
            obj.permission = data.permission;
        } else {
            obj.permission = 'General';
        }
        obj.ipWhitelist = data.ipWhitelist;
        if (data.expire) {
            obj.expire = data.expire;
        } else {
            obj.expire = AddSubAccountApiReq.ExpireEnum._1;
        }
        obj.subName = data.subName;
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
    static fromJson(input: string): AddSubAccountApiReq {
        return this.fromObject(JSON.parse(input));
    }
    /**
     * Create an object from Js Object.
     */
    static fromObject(jsonObject: Object): AddSubAccountApiReq {
        return plainToClassFromExist(new AddSubAccountApiReq(), jsonObject);
    }
}

export namespace AddSubAccountApiReq {
    export enum ExpireEnum {
        /**
         *
         */
        _1 = <any>'-1',
        /**
         *
         */
        _30 = <any>'30',
        /**
         *
         */
        _90 = <any>'90',
        /**
         *
         */
        _180 = <any>'180',
        /**
         *
         */
        _360 = <any>'360',
    }
}

export class AddSubAccountApiReqBuilder {
    constructor(readonly obj: AddSubAccountApiReq) {
        this.obj = obj;
    }
    /**
     * Password(Must contain 7-32 characters. Cannot contain any spaces.)
     */
    setPassphrase(value: string): AddSubAccountApiReqBuilder {
        this.obj.passphrase = value;
        return this;
    }

    /**
     * Remarks(1~24 characters)
     */
    setRemark(value: string): AddSubAccountApiReqBuilder {
        this.obj.remark = value;
        return this;
    }

    /**
     * [Permissions](https://www.kucoin.com/docs-new/doc-338144)(Only General、Spot、Futures、Margin、InnerTransfer(Flex Transfer) permissions can be set, such as \"General, Trade\". The default is \"General\")
     */
    setPermission(value: string): AddSubAccountApiReqBuilder {
        this.obj.permission = value;
        return this;
    }

    /**
     * IP whitelist(You may add up to 20 IPs. Use a halfwidth comma to each IP)
     */
    setIpWhitelist(value: string): AddSubAccountApiReqBuilder {
        this.obj.ipWhitelist = value;
        return this;
    }

    /**
     * API expiration time; Never expire(default)-1，30Day30，90Day90，180Day180，360Day360
     */
    setExpire(value: AddSubAccountApiReq.ExpireEnum): AddSubAccountApiReqBuilder {
        this.obj.expire = value;
        return this;
    }

    /**
     * Sub-account name, create sub account name of API Key.
     */
    setSubName(value: string): AddSubAccountApiReqBuilder {
        this.obj.subName = value;
        return this;
    }

    /**
     * Get the final object.
     */
    build(): AddSubAccountApiReq {
        return this.obj;
    }
}
