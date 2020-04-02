package main

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/model"
	"github.com/ShowPay/ShowMoneyWalletService/util"
	posAPI "github.com/ShowPay/ShowMoneyWalletService/walletAPI"
	"github.com/ShowPay/ShowMoneyWalletService/walletWS"
	"github.com/gookit/ini/v2"
	"net/http"
	"path/filepath"
)

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	model.Wallet_WS_Port = ini.String("Wallet_WS_Port")
}

func startWS() {
	http.HandleFunc("/ws", walletWS.WsHandler)
	fmt.Println("Start showWallet-WS-Service...")
	fmt.Println(util.AddStr("Listen: ws://:", model.Wallet_WS_Port, "/ws"))
	err := http.ListenAndServe(util.AddStr(":", model.Wallet_WS_Port), nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	go posAPI.StartAPI()
	startWS()
}
