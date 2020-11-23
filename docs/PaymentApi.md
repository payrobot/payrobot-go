# \PaymentApi

All URIs are relative to *https://api.payrobot.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentApi.md#CreatePayment) | **Post** /{currency}/payments | Generate a new one-use address to receive a payment
[**GetPayment**](PaymentApi.md#GetPayment) | **Get** /{currency}/payments/{paymentId} | Get detailed information about a payment



## CreatePayment

> PaymentRequest CreatePayment(ctx, currency, type_, destination, amount, callback, optional)

Generate a new one-use address to receive a payment

Generates a new one-use address to receive a payment. It callbacks your web/app server as soon as it's paid and confirmed.  **Payment can be `forwarded` to another address or it can be `stored` in a payrobot.io wallet**     --- ## Important    * Unpaid requests are deleted after **3 hours** of theirs creation   * Confirmed payments information is deleted after **3 days** of being confirmed    --- ## Minimum Amounts     * `Bitcoin`: 0.0001 BTC   * `Litecoin`: 0.001 LTC   * `Bitcoin Cash`: 0.001 BCH    --- ## Callbacks A **payment notificacion** will be sent to your callback url in the following scenarios:    1. *Payment is received partially*   2. *Payment is being confirmed by network*   3. *Payment is confirmed at least with 1 confirmation*   The callback sent to your callback url is a **POST** request with the following parameters:  *Example:*      currency:         \"BTC\"     paymentId:        \"698fd3f6-5482-4798-8a46-6732af440616\"     address:          \"3KoUDMfrov91G4SXaCKGvTWDjGia9Jod5b\"     type:             0     partialAmount:    \"0.00\"                       //Partial amount received when payment is incomplete     remainingAmount:  \"0.00\"                       //Remaining amount to pay when payment is incomplete     amount:           \"0.1\"     feePct:           0.90     feeAmount:        \"0.0009\"     finalAmount:      \"0.0991\"     destination:      \"698fd3f6-5482-4798-8a46-6732af440616\"     reference:        \"12345\"     status:           2 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**type_** | **int32**| * &#x60;0: Receive and forward&#x60; payment is forwarded to a desired coin address once it&#39;s confirmed  * &#x60;1: Receive and store&#x60; payment is stored in a payrobot.io wallet  | 
**destination** | **string**| * For &#x60;Receive and forward&#x60; payment is the &#x60;ADDRESS&#x60; where the payment is going to be forwarded as soon as it&#39;s confirmed. **ADDRESS HAS TO BE OF THE SAME TYPE OF CURRENCY**  * For &#x60;Receive and store&#x60; payment is the payrobot.io &#x60;WALLET ID&#x60; where the payment is going to be stored as soon as it&#39;s confirmed. **WALLET HAS TO BE OF THE SAME TYPE OF CURRENCY**  | 
**amount** | **float32**| Amount of the payment | 
**callback** | **string**| Your URL where payrobot.io will send the status of the payment (Webhook) | 
 **optional** | ***CreatePaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePaymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **reference** | **optional.String**| Optional custom reference to identify the payment | 

### Return type

[**PaymentRequest**](PaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayment

> PaymentRequest GetPayment(ctx, currency, paymentId)

Get detailed information about a payment

Gets detailed information about a payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Object Currency:   * &#x60;btc&#x60;: Bitcoin   * &#x60;ltc&#x60;: Litecoin   * &#x60;bch&#x60;: Bitcoin Cash  | 
**paymentId** | **string**| Payment ID to query | 

### Return type

[**PaymentRequest**](PaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

