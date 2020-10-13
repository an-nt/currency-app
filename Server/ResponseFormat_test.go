package Server

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestFormatResponse(t *testing.T) {
	testIdx := []struct {
		Message      string
		Detail       interface{}
		Code         int
		CodeResponse int
		BodyResponse string
	}{
		{"Success", nil, 200, 200, `{"Message":"Success","Detail":null}`},
		{"Fail", nil, 400, 400, `{"Message":"Fail","Detail":null}`},
	}

	for _, test := range testIdx {
		resp := ResponseFormat{
			Message: test.Message,
			Detail:  test.Detail,
		}

		rec := httptest.NewRecorder()
		resp.FormatResp(rec, test.Code)

		if rec.Code != test.CodeResponse {
			t.Errorf("Wrong status code, have %d, want %d", rec.Code, test.CodeResponse)
		}

		if fmt.Sprintf("%s\n", test.BodyResponse) != fmt.Sprint(rec.Body) {
			t.Errorf("Wrong body content, have %s, want %s", rec.Body, test.BodyResponse)
		}
	}
}
