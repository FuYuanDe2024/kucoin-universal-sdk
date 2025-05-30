// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

import { Transport } from '@internal/interfaces/transport';
import { GetFuturesLedgerReq } from './model_get_futures_ledger_req';
import { GetCrossMarginAccountReq } from './model_get_cross_margin_account_req';
import { GetAccountInfoResp } from './model_get_account_info_resp';
import { GetIsolatedMarginAccountListV1Req } from './model_get_isolated_margin_account_list_v1_req';
import { GetFuturesAccountReq } from './model_get_futures_account_req';
import { GetCrossMarginAccountResp } from './model_get_cross_margin_account_resp';
import { GetSpotLedgerReq } from './model_get_spot_ledger_req';
import { GetSpotAccountListReq } from './model_get_spot_account_list_req';
import { GetSpotAccountDetailReq } from './model_get_spot_account_detail_req';
import { GetMarginHFLedgerResp } from './model_get_margin_hf_ledger_resp';
import { GetFuturesAccountResp } from './model_get_futures_account_resp';
import { GetIsolatedMarginAccountReq } from './model_get_isolated_margin_account_req';
import { GetIsolatedMarginAccountDetailV1Req } from './model_get_isolated_margin_account_detail_v1_req';
import { GetSpotHFLedgerResp } from './model_get_spot_hf_ledger_resp';
import { GetSpotLedgerResp } from './model_get_spot_ledger_resp';
import { GetIsolatedMarginAccountDetailV1Resp } from './model_get_isolated_margin_account_detail_v1_resp';
import { GetIsolatedMarginAccountResp } from './model_get_isolated_margin_account_resp';
import { GetSpotHFLedgerReq } from './model_get_spot_hf_ledger_req';
import { GetFuturesLedgerResp } from './model_get_futures_ledger_resp';
import { GetApikeyInfoResp } from './model_get_apikey_info_resp';
import { GetSpotAccountListResp } from './model_get_spot_account_list_resp';
import { GetMarginAccountDetailResp } from './model_get_margin_account_detail_resp';
import { GetIsolatedMarginAccountListV1Resp } from './model_get_isolated_margin_account_list_v1_resp';
import { GetMarginHFLedgerReq } from './model_get_margin_hf_ledger_req';
import { GetSpotAccountDetailResp } from './model_get_spot_account_detail_resp';
import { GetSpotAccountTypeResp } from './model_get_spot_account_type_resp';

export interface AccountAPI {
    /**
     * getAccountInfo Get Account Summary Info
     * Description: This endpoint can be used to obtain account summary information.
     * Documentation: https://www.kucoin.com/docs-new/api-3470119
     * +---------------------+------------+
     * | Extra API Info      | Value      |
     * +---------------------+------------+
     * | API-DOMAIN          | SPOT       |
     * | API-CHANNEL         | PRIVATE    |
     * | API-PERMISSION      | GENERAL    |
     * | API-RATE-LIMIT-POOL | MANAGEMENT |
     * | API-RATE-LIMIT      | 20         |
     * +---------------------+------------+
     */
    getAccountInfo(): Promise<GetAccountInfoResp>;

    /**
     * getApikeyInfo Get Apikey Info
     * Description: Get the information of the api key. Use the api key pending to be checked to call the endpoint. Both master and sub user\&#39;s api key are applicable.
     * Documentation: https://www.kucoin.com/docs-new/api-3470130
     * +---------------------+------------+
     * | Extra API Info      | Value      |
     * +---------------------+------------+
     * | API-DOMAIN          | SPOT       |
     * | API-CHANNEL         | PRIVATE    |
     * | API-PERMISSION      | GENERAL    |
     * | API-RATE-LIMIT-POOL | MANAGEMENT |
     * | API-RATE-LIMIT      | 20         |
     * +---------------------+------------+
     */
    getApikeyInfo(): Promise<GetApikeyInfoResp>;

    /**
     * getSpotAccountType Get Account Type - Spot
     * Description: This interface determines whether the current user is a spot high-frequency user or a spot low-frequency user.
     * Documentation: https://www.kucoin.com/docs-new/api-3470120
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 30      |
     * +---------------------+---------+
     */
    getSpotAccountType(): Promise<GetSpotAccountTypeResp>;

    /**
     * getSpotAccountList Get Account List - Spot
     * Description: Get a list of accounts. Please Deposit to the main account firstly, then transfer the funds to the trade account via Inner Transfer before transaction.
     * Documentation: https://www.kucoin.com/docs-new/api-3470125
     * +---------------------+------------+
     * | Extra API Info      | Value      |
     * +---------------------+------------+
     * | API-DOMAIN          | SPOT       |
     * | API-CHANNEL         | PRIVATE    |
     * | API-PERMISSION      | GENERAL    |
     * | API-RATE-LIMIT-POOL | MANAGEMENT |
     * | API-RATE-LIMIT      | 5          |
     * +---------------------+------------+
     */
    getSpotAccountList(req: GetSpotAccountListReq): Promise<GetSpotAccountListResp>;

    /**
     * getSpotAccountDetail Get Account Detail - Spot
     * Description: get Information for a single spot account. Use this endpoint when you know the accountId.
     * Documentation: https://www.kucoin.com/docs-new/api-3470126
     * +---------------------+------------+
     * | Extra API Info      | Value      |
     * +---------------------+------------+
     * | API-DOMAIN          | SPOT       |
     * | API-CHANNEL         | PRIVATE    |
     * | API-PERMISSION      | GENERAL    |
     * | API-RATE-LIMIT-POOL | MANAGEMENT |
     * | API-RATE-LIMIT      | 5          |
     * +---------------------+------------+
     */
    getSpotAccountDetail(req: GetSpotAccountDetailReq): Promise<GetSpotAccountDetailResp>;

    /**
     * getCrossMarginAccount Get Account - Cross Margin
     * Description: Request via this endpoint to get the info of the cross margin account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470127
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 15      |
     * +---------------------+---------+
     */
    getCrossMarginAccount(req: GetCrossMarginAccountReq): Promise<GetCrossMarginAccountResp>;

    /**
     * getIsolatedMarginAccount Get Account - Isolated Margin
     * Description: Request via this endpoint to get the info of the isolated margin account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470128
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 15      |
     * +---------------------+---------+
     */
    getIsolatedMarginAccount(
        req: GetIsolatedMarginAccountReq,
    ): Promise<GetIsolatedMarginAccountResp>;

    /**
     * getFuturesAccount Get Account - Futures
     * Description: Request via this endpoint to get the info of the futures account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470129
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | FUTURES |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | FUTURES |
     * | API-RATE-LIMIT-POOL | FUTURES |
     * | API-RATE-LIMIT      | 5       |
     * +---------------------+---------+
     */
    getFuturesAccount(req: GetFuturesAccountReq): Promise<GetFuturesAccountResp>;

    /**
     * getSpotLedger Get Account Ledgers - Spot/Margin
     * Description: This interface is for transaction records from all types of your accounts, supporting inquiry of various currencies. Items are paginated and sorted to show the latest first. See the Pagination section for retrieving additional entries after the first page.
     * Documentation: https://www.kucoin.com/docs-new/api-3470121
     * +---------------------+------------+
     * | Extra API Info      | Value      |
     * +---------------------+------------+
     * | API-DOMAIN          | SPOT       |
     * | API-CHANNEL         | PRIVATE    |
     * | API-PERMISSION      | GENERAL    |
     * | API-RATE-LIMIT-POOL | MANAGEMENT |
     * | API-RATE-LIMIT      | 2          |
     * +---------------------+------------+
     */
    getSpotLedger(req: GetSpotLedgerReq): Promise<GetSpotLedgerResp>;

    /**
     * getSpotHFLedger Get Account Ledgers - Trade_hf
     * Description: This API endpoint returns all transfer (in and out) records in high-frequency trading account and supports multi-coin queries. The query results are sorted in descending order by createdAt and id.
     * Documentation: https://www.kucoin.com/docs-new/api-3470122
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 2       |
     * +---------------------+---------+
     */
    getSpotHFLedger(req: GetSpotHFLedgerReq): Promise<GetSpotHFLedgerResp>;

    /**
     * getMarginHFLedger Get Account Ledgers - Margin_hf
     * Description: This API endpoint returns all transfer (in and out) records in high-frequency margin trading account and supports multi-coin queries. The query results are sorted in descending order by createdAt and id.
     * Documentation: https://www.kucoin.com/docs-new/api-3470123
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 2       |
     * +---------------------+---------+
     */
    getMarginHFLedger(req: GetMarginHFLedgerReq): Promise<GetMarginHFLedgerResp>;

    /**
     * getFuturesLedger Get Account Ledgers - Futures
     * Description: This interface can query the ledger records of the futures business line
     * Documentation: https://www.kucoin.com/docs-new/api-3470124
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | FUTURES |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | FUTURES |
     * | API-RATE-LIMIT      | 2       |
     * +---------------------+---------+
     */
    getFuturesLedger(req: GetFuturesLedgerReq): Promise<GetFuturesLedgerResp>;

    /**
     * @deprecated
     * getMarginAccountDetail Get Account Detail - Margin
     * Description: Request via this endpoint to get the info of the margin account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470311
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 40      |
     * +---------------------+---------+
     */
    getMarginAccountDetail(): Promise<GetMarginAccountDetailResp>;

    /**
     * @deprecated
     * getIsolatedMarginAccountListV1 Get Account List - Isolated Margin - V1
     * Description: Request via this endpoint to get the info list of the isolated margin account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470314
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 50      |
     * +---------------------+---------+
     */
    getIsolatedMarginAccountListV1(
        req: GetIsolatedMarginAccountListV1Req,
    ): Promise<GetIsolatedMarginAccountListV1Resp>;

    /**
     * @deprecated
     * getIsolatedMarginAccountDetailV1 Get Account Detail - Isolated Margin - V1
     * Description: Request via this endpoint to get the info of the isolated margin account.
     * Documentation: https://www.kucoin.com/docs-new/api-3470315
     * +---------------------+---------+
     * | Extra API Info      | Value   |
     * +---------------------+---------+
     * | API-DOMAIN          | SPOT    |
     * | API-CHANNEL         | PRIVATE |
     * | API-PERMISSION      | GENERAL |
     * | API-RATE-LIMIT-POOL | SPOT    |
     * | API-RATE-LIMIT      | 50      |
     * +---------------------+---------+
     */
    getIsolatedMarginAccountDetailV1(
        req: GetIsolatedMarginAccountDetailV1Req,
    ): Promise<GetIsolatedMarginAccountDetailV1Resp>;
}

export class AccountAPIImpl implements AccountAPI {
    constructor(private transport: Transport) {}

    getAccountInfo(): Promise<GetAccountInfoResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v2/user-info',
            null,
            GetAccountInfoResp,
            false,
        );
    }

    getApikeyInfo(): Promise<GetApikeyInfoResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/user/api-key',
            null,
            GetApikeyInfoResp,
            false,
        );
    }

    getSpotAccountType(): Promise<GetSpotAccountTypeResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/hf/accounts/opened',
            null,
            GetSpotAccountTypeResp,
            false,
        );
    }

    getSpotAccountList(req: GetSpotAccountListReq): Promise<GetSpotAccountListResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/accounts',
            req,
            GetSpotAccountListResp,
            false,
        );
    }

    getSpotAccountDetail(req: GetSpotAccountDetailReq): Promise<GetSpotAccountDetailResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/accounts/{accountId}',
            req,
            GetSpotAccountDetailResp,
            false,
        );
    }

    getCrossMarginAccount(req: GetCrossMarginAccountReq): Promise<GetCrossMarginAccountResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v3/margin/accounts',
            req,
            GetCrossMarginAccountResp,
            false,
        );
    }

    getIsolatedMarginAccount(
        req: GetIsolatedMarginAccountReq,
    ): Promise<GetIsolatedMarginAccountResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v3/isolated/accounts',
            req,
            GetIsolatedMarginAccountResp,
            false,
        );
    }

    getFuturesAccount(req: GetFuturesAccountReq): Promise<GetFuturesAccountResp> {
        return this.transport.call(
            'futures',
            false,
            'GET',
            '/api/v1/account-overview',
            req,
            GetFuturesAccountResp,
            false,
        );
    }

    getSpotLedger(req: GetSpotLedgerReq): Promise<GetSpotLedgerResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/accounts/ledgers',
            req,
            GetSpotLedgerResp,
            false,
        );
    }

    getSpotHFLedger(req: GetSpotHFLedgerReq): Promise<GetSpotHFLedgerResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/hf/accounts/ledgers',
            req,
            GetSpotHFLedgerResp,
            false,
        );
    }

    getMarginHFLedger(req: GetMarginHFLedgerReq): Promise<GetMarginHFLedgerResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v3/hf/margin/account/ledgers',
            req,
            GetMarginHFLedgerResp,
            false,
        );
    }

    getFuturesLedger(req: GetFuturesLedgerReq): Promise<GetFuturesLedgerResp> {
        return this.transport.call(
            'futures',
            false,
            'GET',
            '/api/v1/transaction-history',
            req,
            GetFuturesLedgerResp,
            false,
        );
    }

    getMarginAccountDetail(): Promise<GetMarginAccountDetailResp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/margin/account',
            null,
            GetMarginAccountDetailResp,
            false,
        );
    }

    getIsolatedMarginAccountListV1(
        req: GetIsolatedMarginAccountListV1Req,
    ): Promise<GetIsolatedMarginAccountListV1Resp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/isolated/accounts',
            req,
            GetIsolatedMarginAccountListV1Resp,
            false,
        );
    }

    getIsolatedMarginAccountDetailV1(
        req: GetIsolatedMarginAccountDetailV1Req,
    ): Promise<GetIsolatedMarginAccountDetailV1Resp> {
        return this.transport.call(
            'spot',
            false,
            'GET',
            '/api/v1/isolated/account/{symbol}',
            req,
            GetIsolatedMarginAccountDetailV1Resp,
            false,
        );
    }
}
