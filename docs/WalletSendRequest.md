# WalletSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**CryptoCurrency**](CryptoCurrency.md) |  | [optional] 
**WalletId** | **string** | Unique ID of the new created Wallet | [optional] 
**RequestId** | **string** | ID of this transaction, it can be used letter in the callback or to query it | [optional] 
**Timestamp** | **int32** | Request creation date expressed in UNIX timestamp | [optional] 
**Lastupdate** | **int32** | Last update expressed in UNIX timestamp | [optional] 
**Amount** | **string** | Total amount sent from wallet | [optional] 
**Callback** | **string** | Optional callback to notify your web / app as soon as the send request has been fully broadcasted to the network | [optional] 
**Destination** | [**[]AddressDetail**](AddressDetail.md) | Array with all the destination coin addres(es) and the amount(s) to send  | [optional] 
**Txid** | **string** | Tx Hash. *Only available in requests with confirmed status*  | [optional] 
**Status** | **int32** | Status of this send request:   * &#x60;0: Queued&#x60; Request has been queued for broadcasting. It usually takes few seconds under normal conditions   * &#x60;1: Broadcasted&#x60; Request has been fully broadcasted to Bitcoin Network   | [optional] [default to 0]
**Error** | **bool** | &#x60;true&#x60; is there was a problem. &#x60;false&#x60; if not  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


