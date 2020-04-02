package posAPI

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/model"
	"github.com/ShowPay/ShowMoneyWalletService/util"
	"github.com/gookit/ini/v2"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	model.Wallet_API_Port = ini.String("Wallet_API_Port")
}

func StartAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet/notifyPaymentRequest", NotifyPaymentRequest).Methods("POST")
	fmt.Println("Start showWallet-API-Service...")
	err := http.ListenAndServe(util.AddStr(":", model.Wallet_API_Port), r)
	if err != nil {
		panic(err)
	}
}