package posAPI

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyWalletService/util"
	"github.com/gookit/ini/v2"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

var wallet_api_port = "1234"

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	wallet_api_port = ini.String("wallet_api_port")
}

func StartAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet/notifyPaymentRequest", NotifyPaymentRequest).Methods("POST")
	fmt.Println("Start showWallet-API-Service...")
	err := http.ListenAndServe(util.AddStr(":", wallet_api_port), r)
	if err != nil {
		panic(err)
	}
}