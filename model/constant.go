package model


var (
	Wallet_API_Port = ""
	Wallet_WS_Port = ""
	Domain_Wallet_WS = "localhost:" + Wallet_WS_Port

)

const (
	//HEART
	HEART_BEAT = "HEART_BEAT"

	//wallet
	WS_WALLET_CONNECT = "WS_WALLET_CONNECT"
	WS_API_NOTIFY_TO_WALLET = "WS_API_NOTIFY"
	WS_WALLET_NOTIFY = "WS_WALLET_NOTIFY"

	//comment
	WS_RESPONSE_SUCCESS = "WS_RESPONSE_SUCCESS"
	WS_RESPONSE_ERROR = "WS_RESPONSE_ERROR"

	//disconnect
	WS_DISCONNECT = "WS_DISCONNECT"
)

const (
	WS_CODE_SERVER = 0 //Proactive notification
	WS_CODE_SEND_SUCCESS = 200
	WS_CODE_SEND_ERROR = 400
	WS_CODE_HEART_BEAT = 10

)