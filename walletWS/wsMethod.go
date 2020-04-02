package walletWS

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/model"
	"github.com/ShowPay/ShowMoneyWalletService/service"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

//Push message heartbeat back
func SendHeartBeat(c net.Conn) {
	wsData := &model.WsData{
		M: model.HEART_BEAT,
		C: model.WS_CODE_HEART_BEAT,
		D: "Heart beat back",
	}
	SendMsgToConn(c, wsData)
}

//cache wallet conn
func GetWalletId(c net.Conn, wsItemMap *WsItemMap, msg string) error {
	req := &model.WsWalletReq{}
	err := new(service.Ws_showWallet).MsgToWSWallet(msg, req)
	if err != nil {
		fmt.Println("MsgToWSPaymentRequest err:", err)
		return err
	}
	//Caching conn
	wsItemMap.SetConn(req.WalletId, c)
	wsData := &model.WsData{
		M: model.WS_RESPONSE_SUCCESS,
		C: model.WS_CODE_SEND_SUCCESS,
		D: "connect success",
	}
	SendMsgToConn(c, wsData)
	return nil
}


//API send PaymentRequest to wallet
func NotifyPRToWallet(c net.Conn, wsItemMap *WsItemMap, msg string) {
	req := &model.ApiPaymentRequestReq{}
	err := new(service.Ws_showWallet).MsgToApiPaymentRequest(msg, req)
	if err != nil {
		fmt.Println("MsgToApiPaymentRequest err:", err)
		return
	}

	//get wallet conn
	if c, ok := wsItemMap.GetConn(req.WalletId); ok {
		wsData := &model.WsData{
			M: model.WS_WALLET_NOTIFY,
			C: model.WS_CODE_SERVER,
			D: req,
		}
		SendMsgToConn(c, wsData)
		return
	}
	fmt.Println("wallet is not exist | id: [", req.WalletId, "]")
}

//WS Reply
func SendWSResponseForErr(c net.Conn, wsItemMap *WsItemMap, msg string) {
	if _, ok := wsItemMap.Get(c); !ok {
		fmt.Println("conn is not exist")
		return
	} else {
		wsData := &model.WsData{
			M: model.WS_RESPONSE_ERROR,
			C: model.WS_CODE_SEND_ERROR,
			D: msg,
		}
		SendMsgToConn(c, wsData)
		return
	}
}

func SendMsgToConn(c net.Conn, wsData *model.WsData)  {
	wsDataStr,err := wsData.ToString()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = wsutil.WriteServerMessage(c, ws.OpText, []byte(wsDataStr))
	if err != nil {
		// handle error
		fmt.Println("WriteServerMessage Err:" + err.Error())
	}
	return
}