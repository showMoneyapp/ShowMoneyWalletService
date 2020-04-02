package main

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/util"
	posAPI "github.com/ShowPay/ShowMoneyWalletService/walletAPI"
	"github.com/ShowPay/ShowMoneyWalletService/walletWS"
	"net/http"
	"path/filepath"
)

var wallet_ws_port = "1234"

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	wallet_ws_port = ini.String("wallet_ws_port")
}

func startWS() {
	http.HandleFunc("/ws", walletWS.WsHandler)
	fmt.Println("Start showWallet-WS-Service...")
	err := http.ListenAndServe(util.AddStr(":", wallet_ws_port), nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	go posAPI.StartAPI()
	startWS()
}
