# \WalletApi

All URIs are relative to *https://api.payrobot.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWallet**](WalletApi.md#CreateWallet) | **Post** /{currency}/wallets | Create new wallet
[**CreateWalletSendRequest**](WalletApi.md#CreateWalletSendRequest) | **Post** /{currency}/wallets/{walletId}/send-requests | Send funds from a wallet
[**GetWallet**](WalletApi.md#GetWallet) | **Get** /{currency}/wallets/{walletId} | Get Wallet information
[**GetWalletHistory**](WalletApi.md#GetWalletHistory) | **Get** /{currency}/wallets/{walletId}/history | Get last transactions of wallet
[**GetWalletSendRequest**](WalletApi.md#GetWalletSendRequest) | **Get** /{currency}/wallets/{walletId}/send-requests/{requestId} | Obtain information of a send request



## CreateWallet

> WalletCreationInfo CreateWallet(ctx, currency)

Create new wallet

Creates a new wallet where you can receive, store and send funds for your web or app.  --- ## Important This method returns your `Wallet Passphrase`, it will be required when you send funds from your wallet. **Please keep it safe and private** 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 

### Return type

[**WalletCreationInfo**](WalletCreationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWalletSendRequest

> WalletSendRequest CreateWalletSendRequest(ctx, currency, walletId, destination, seed, token, optional)

Sends funds from a wallet to one or multiple addresses.


---

## Required Authorization Token

This transaction requires an authorization `token` which is the result
of the `sha-256` hash of the following string:



    walletId~destination~seed~walletPassphrase



**For example**


Considering the following example values for the token:
  - `walletId` 9df3f909-088d-4724-b34f-9a587c5ccc15
  - `destination`
    [{"address":"bc1q5defveu0acrf87m3huwxjq6pqaszdjf3d4ej9y","amount":0.01},{"address":"bc1qs59a7e23zpjm0znteytrxvj839dlp205e50zch","amount":0.056}]


  - `seed` 758748394
  - `walletPassphrase` **Note: this was provided when you created the wallet** OHh6IIININmfmjGGsxlBBft2ch61VncaPscsp295h2ULx9xPY07Jom3d5cBifgoW

The resulting string, previous to hash is::



    9df3f909-088d-4724-b34f-9a587c5ccc15~[{"address":"bc1q5defveu0acrf87m3huwxjq6pqaszdjf3d4ej9y","amount":0.01},{"address":"bc1qs59a7e23zpjm0znteytrxvj839dlp205e50zch","amount":0.056}]~758748394~OHh6IIININmfmjGGsxlBBft2ch61VncaPscsp295h2ULx9xPY07Jom3d5cBifgoW



Finally after applying `sha-256` hash, we obtain the required `token`:



    804ca9457b0fe3e4d243fe9e39e760ff1f287491ae8e79d015f92f7c6c96d7b1


---

## Important

  * Send requests are commonly queued, optionally you can specify a callback to get your web / app notified as soon as the request has been fully broadcasted to the Network.

  * Transaction is limited to `25` destination addresses per request

  * Tx Hash is provided only through the callback

  * Confirmed send requests information is `DELETED`
after `3 days` of being confirmed

---

## Minimum Send Amounts


  * `Bitcoin`: 0.0001 BTC
  * `Litecoin`: 0.001 LTC
  * `Bitcoin Cash`: 0.001 BCH

---

## Callback

Send requests are commonly queued, optionally you can specify a callback
to get your web / app notified as soon as the request has been fully
broadcasted to the Network.


The callback sent to your callback url is a **POST** request with the
following parameters:


*Example:*

    currency:     "BTC"
    walletId:     "698fd3f6-5482-4798-8a46-6732af440616"
    requestId:    "123fd3f6-9078-5790-4f40-6932bf440120"
    timestamp:    1577179288
    lastupdate:   1577179388
    amount:       "0.01"
    callback:     "https://callback-url.com"
    destination:  '[{"address": "bc1qf6ss0qtdn5q42..."
                  "amount": "0.01"}]'
    txid:         "2cdac43e92e65cb428e3ed992bcf61347..."
    status:       0

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**walletId** | **string**| Wallet where funds to send are stored | 
**destination** | **string**| JSON formatted array with all the destination addres(es) and the amount(s) to send\\ &#x60;[{\&quot;address\&quot;:\&quot;desired-destination-address\&quot;,\&quot;amount\&quot;:X.XXXXXXXX}, ...]&#x60;  | 
**seed** | **string**| Unique random string generated by your web/app. **IT MUST BE UNIQUE PER TRANSACTION PER WALLET** | 
**token** | **string**| SHA-256 hash of the concatenated string (substituting with the proper data):\\ &#x60;walletId~destination~seed~walletPassphrase&#x60;  | 
 **optional** | ***CreateWalletSendRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWalletSendRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **callback** | **optional.String**| Optional callback to notify your web / app as soon as the send request has been fully broadcasted to the Network | 

### Return type

[**WalletSendRequest**](WalletSendRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> Wallet GetWallet(ctx, currency, walletId)

Get Wallet information

Gets detailed information from a Wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**walletId** | **string**| ID of the desired Wallet | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletHistory

> WalletHistory GetWalletHistory(ctx, currency, walletId)

Get last transactions of wallet

Gets last 100 transactions of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**walletId** | **string**| ID of the desired Wallet | 

### Return type

[**WalletHistory**](WalletHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletSendRequest

> WalletSendRequest GetWalletSendRequest(ctx, currency, walletId, requestId)

Obtain information of a send request

Obtains detailed information about a send request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**walletId** | **string**| Wallet where funds to send are stored | 
**requestId** | **string**| Send Request ID | 

### Return type

[**WalletSendRequest**](WalletSendRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

