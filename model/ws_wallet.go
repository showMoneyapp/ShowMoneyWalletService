package model

type WsWalletReq struct {
	FeatureCode string `json:"featureCode"` //"show&pay"
	WalletId    string `json:"walletId"`    //string.  required. wallet-device
}
