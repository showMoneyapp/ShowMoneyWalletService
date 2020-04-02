package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/model"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type Ws_showWallet struct {}

//Parsing WSWallet
func (sw *Ws_showWallet) MsgToWSWallet(msg string, req *model.WsWalletReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("Json data parsed failed.")
	}
	return nil
}

//Parsing ApiPaymentRequest
func (sw *Ws_showWallet) MsgToApiPaymentRequest(msg string, req *model.ApiPaymentRequestReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("Json data parsed failed.")
	}
	return nil
}

//Forward notify wallet. ws notification
func (sw *Ws_showWallet) NotifyPaymentRequest(req *model.ApiPaymentRequestReq, resp *model.ApiPaymentRequestResp) error {
	if req.Outputs == nil || len(req.Outputs) == 0 {
		return errors.New("outputs is empty.")
	}
	if req.Network == "" || len(req.Network) == 0 {
		return errors.New("Network is empty.")
	}
	if req.WalletId == "" || len(req.WalletId) == 0 {
		return errors.New("WalletId is empty.")
	}
	if req.DeviceId == "" || len(req.DeviceId) == 0 {
		return errors.New("DeviceId is empty.")
	}
	if req.PaymentUrl == "" || len(req.PaymentUrl) == 0 {
		return errors.New("PaymentUrl is empty.")
	}
	if req.CreationTimestamp  == 0 {
		return errors.New("CreationTimestamp is empty.")
	}
	if req.ExpirationTimestamp  == 0 {
		return errors.New("ExpirationTimestamp is empty.")
	}
	for _, v := range req.Outputs {
		if v.Amount < 0 || v.Script == "" || len(v.Script) == 0 {
			return errors.New("output.Amount or Script is empty.")
		}
	}


	u := url.URL{Scheme:"ws", Host:model.Domain_wallet_ws, Path:"/ws"}
	fmt.Println("connecting to :", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("dial:" + err.Error())
	}


	//Forward PaymentRequest
	wsData := &model.WsData{
		M: model.WS_API_NOTIFY_TO_WALLET,
		C: 0,
		D: req,
	}
	wsDataStr, err := wsData.ToString()
	if err != nil {
		return errors.New("wsData conversion failed:" + err.Error())
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(wsDataStr))
	if err != nil {
		log.Println("write | err :", err)
		return errors.New("write | err :" + err.Error())
	}

	//disconnect
	wsDataDis := &model.WsData{
		M: model.WS_DISCONNECT,
		C: 0,
		D: "disconnect",
	}
	wsDataDisStr, err := wsDataDis.ToString()
	if err != nil {
		return errors.New("wsDataDis conversion failed:" + err.Error())
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(wsDataDisStr))
	if err != nil {
		log.Println("write | err :", err)
		return errors.New("write | err :" + err.Error())
	}

	time.Sleep(100*time.Millisecond)
	defer c.Close()

	resp.Code = 200
	resp.Message = "Notify Success"
	return nil
}