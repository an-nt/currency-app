package Server

import (
	"encoding/json"
	"net/http"
)

type FmResp struct {
	Message string      `json: "message"`
	Detail  interface{} `json: "detail"`
}

type IFmResp interface {
	FormatResp(w http.ResponseWriter, status int)
}

func NewFmResp(message string, detail interface{}) IFmResp {
	return &FmResp{
		Message: message,
		Detail:  detail,
	}
}

func (f *FmResp) FormatResp(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "JSON")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(f)
	if err != nil {
		resp := NewFmResp("Internal Server Error", nil)
		resp.FormatResp(w, http.StatusInternalServerError)
	}
}
