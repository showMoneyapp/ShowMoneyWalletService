package model

type ApiPaymentRequestReq struct {
	Network             string       `json:"network"`             //string. required. bitcoin network tag：Always set to "bitcoin".
	Outputs             []*ApiOutput `json:"outputs"`             //an array of outputs. required, but can have zero elements.
	WalletId            string       `json:"walletId"`            //string. required. wallet-device
	DeviceId            string       `json:"deviceId"`            //string. required. pos-device
	CreationTimestamp   int64        `json:"creationTimestamp"`   //number. required.
	ExpirationTimestamp int64        `json:"expirationTimestamp"` //number. optional.
	Memo                string      `json:"memo"`                //string. optional.
	PaymentUrl          string      `json:"paymentUrl"`          //string. required.
	MerchantData        string      `json:"merchantData"`        //string. optional.
	QRCodeLabelData     string      `json:"qrcodeLabelData"`     //string. optional.
}

type ApiPaymentRequestResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
}

type ApiOutput struct {
	Amount      int64  `json:"amount"`      //satoshi. number. required.
	Script      string `json:"script"`      //string. required. hexadecimal script.
	Description string `json:"description"` //string. optional. must not have JSON string length of greater than 100.
	TokenIndex  int64  `json:"tokenIndex"`  //number. token_index. default 0, 0 mean bsv
}
