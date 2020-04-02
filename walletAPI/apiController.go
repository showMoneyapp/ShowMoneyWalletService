package posAPI

import (
	"github.com/ShowPay/ShowMoneyWalletService/model"
	"github.com/ShowPay/ShowMoneyWalletService/service"
	"net/http"
)

//Notify PaymentRequest
func NotifyPaymentRequest(w http.ResponseWriter, r *http.Request)  {
	req := &model.ApiPaymentRequestReq{}
	resp := &model.ApiPaymentRequestResp{}
	err := pares(r, req)
	if err != nil {
		ResponseError(w, err.Error())
		return
	}
	if err = new(service.Ws_showWallet).NotifyPaymentRequest(req, resp);err != nil{
		ResponseError(w, err.Error())
		return
	}
	ResponseSuccess(w, resp)
}