package Server

import (
	"CurrencyApp/API"
	"net/http"
)

type GetExRateFormatter struct {
	Handler API.IExecGetExRate
}

type IGetExRateFormatter interface {
	FormatGetExRate(w http.ResponseWriter, r *http.Request)
}

func (f *GetExRateFormatter) FormatGetExRate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	rate, err := f.Handler.GetExRate(token)
	if err != nil {
		resp := FmResp{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusForbidden)
		return
	}
	resp := FmResp{
		Message: "Success",
		Detail:  rate,
	}
	resp.FormatResp(w, http.StatusOK)
}
