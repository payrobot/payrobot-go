/*
 * Payrobot
 *
 * # Introduction Accept, store, send or forward Bitcoin, Litecoin and Bitcoin Cash for your website or app and protect your privacy.  Supported crytocurrencies:   * BTC Bitcoin   * LTC Litecoin   * BCH Bitcoin Cash   ## Benefits    * **Anonymous** No personal details are required and transactions are mixed among all payments. You can forward your payments so as soon payrobot.io receives it forwards it to another address under your control.      * **No Registration** No registration, sign-up, application or form required to use payrobot.io      * **Easy Integration** Integrate your web / app through our simple RESTful API, you can accept payments with just one line of code!      * **Instant Payment Notification** Our servers notify your web / app the status of your payments. No polling, daemons or cronjobs required on your side!      * **Secure** Payrobot.io works with SSL and bank-level security protocols. Your transactions are safe!   ## Features **Payment Forward** Generate one-time addresses to recieve payments. Payrobot will notify your web /app through callbacks (webhooks) the status of the payment. As soon as it's confirmed the payment is forwarded to your desired address.  **Wallet** Receive, send payments and store your coins in a secure, private and anonymous wallet. All events are notified to your web / app through callbacks (webhooks). You can generate wallets with just one line of code without registration or further information  ## Fees **Only 0.90% per inbound transaction** (receive payments), NO HIDDEN FEES. All outbound transactions (send funds) are totally free.  Minimum fees applies, therefore the largest amount is going to be considered as fee either: `(inboundAmount*feePct)` or `the minimum fee`  **Inbound Fees (Receive payments)**    - `Bitcoin` 0.90% *(Minimum fee 0.00005 BTC)*   - `Litecoin` 0.90% *(Minimum fee 0.0005 LTC)*   - `Bitcoin Cash` 0.90% *(Minimum fee 0.0005 BCH)*     **Outbound Fees (Send funds)**    - `Bitcoin` 0.00%   - `Litecoin` 0.00%   - `Bitcoin Cash` 0.00%   ## Rate Limit To guarantee the good performance of the service and its fair use. The API is **limited to receiving 120 requests per minute per IP**, which is sufficient for most use cases.  Payrobot.io is asynchronous in most API methods to communicate with your application through callbacks (webhooks), thus reducing unnecessary calls to the service.  **If the limit is exceeded, the IP will be banned for 1 minute.**  If you require an upper limit for your application, do not hesitate to contact us  ## Considerations    * Amounts in responses are expresed as `strings`      * Wallets are not multi-currency, you have to create a different wallet per cryptocurrency (You can't store Litecoin in a Bitcoin wallet and vice-versa)      * Payment forwarding has to be of the same type of currency (You can't forward a Bitcoin Cash payment to a Bitcoin address and vice-versa)    
 *
 * API version: 1.0
 * Contact: contact@payrobot.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payrobot
// WalletSendRequest struct for WalletSendRequest
type WalletSendRequest struct {
	Currency CryptoCurrency `json:"currency,omitempty"`
	// Unique ID of the new created Wallet
	WalletId string `json:"walletId,omitempty"`
	// ID of this transaction, it can be used letter in the callback or to query it
	RequestId string `json:"requestId,omitempty"`
	// Request creation date expressed in UNIX timestamp
	Timestamp int32 `json:"timestamp,omitempty"`
	// Last update expressed in UNIX timestamp
	Lastupdate int32 `json:"lastupdate,omitempty"`
	// Total amount sent from wallet
	Amount string `json:"amount,omitempty"`
	// Optional callback to notify your web / app as soon as the send request has been fully broadcasted to the network
	Callback string `json:"callback,omitempty"`
	// Array with all the destination coin addres(es) and the amount(s) to send 
	Destination []AddressDetail `json:"destination,omitempty"`
	// Tx Hash. *Only available in requests with confirmed status* 
	Txid string `json:"txid,omitempty"`
	// Status of this send request:   * `0: Queued` Request has been queued for broadcasting. It usually takes few seconds under normal conditions   * `1: Broadcasted` Request has been fully broadcasted to Bitcoin Network  
	Status int32 `json:"status,omitempty"`
	// `true` is there was a problem. `false` if not 
	Error bool `json:"error,omitempty"`
}
