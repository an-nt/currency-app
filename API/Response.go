package API

import (
	"encoding/json"
	"net/http"
)

type FmResp struct {
	Statuscode int         `json: "statuscode"`
	Message    string      `json: "message"`
	Detail     interface{} `json: "detail"`
}

type IFmResp interface {
	FormatResp(w http.ResponseWriter)
}

func NewFmResp(statuscode int, message string, detail interface{}) IFmResp {
	return &FmResp{
		Statuscode: statuscode,
		Message:    message,
		Detail:     detail,
	}
}

func (f *FmResp) FormatResp(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "JSON")
	w.WriteHeader(f.Statuscode)

	err := json.NewEncoder(w).Encode(f)
	if err != nil {
		resp := NewFmResp(http.StatusInternalServerError, "Internal Server Error", nil)
		resp.FormatResp(w)
	}
}
