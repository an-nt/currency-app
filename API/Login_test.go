package API

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tcIndex = []struct {
	Mcode   int
	Mmess   string
	Mdetail interface{}
	code    int
	body    string
}{
	{200, "hello", nil, 200, `{"Statuscode":200,"Message":"hello","Detail":null}`},
	{200, "Authorized", "This is a token", 200, `{"Statuscode":200,"Message":"Authorized","Detail":"This is a token"}`},
}

type MockFormatter struct {
	Loop int
}
type IMockFormatter interface {
	FmLogin(*http.Request) (FmResp, error)
}

func (m *MockFormatter) FmLogin(r *http.Request) (FmResp, error) {
	resp := FmResp{
		Statuscode: tcIndex[m.Loop].Mcode,
		Message:    tcIndex[m.Loop].Mmess,
		Detail:     tcIndex[m.Loop].Mdetail,
	}
	return resp, nil
}

func TestLogin(t *testing.T) {
	var api Api
	for i, tc := range tcIndex {
		api.apiLogin = &MockFormatter{Loop: i}

		req, err := http.NewRequest("POST", "http://localhost:8080/v1/login", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		api.Login(rec, req)

		if rec.Code != tc.code {
			t.Errorf("Wrong status code, have %v, want %v", rec.Code, tc.code)
		}

		expBody := fmt.Sprintf("%s\n", tc.body)
		if rec.Body.String() != expBody {
			t.Errorf("Wrong body content, have %s, want %s", rec.Body.String(), tc.body)
		}
	}
}
