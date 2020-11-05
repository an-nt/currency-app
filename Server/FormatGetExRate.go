package Server

import (
	"net/http"
)

func (sv *HttpServer) formatGetExRate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	rate, err := sv.Exec.GetExRate(token)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusForbidden)
		return
	}
	resp := ResponseFormat{
		Message: "Success",
		Detail:  rate,
	}
	resp.FormatResp(w, http.StatusOK)
}
