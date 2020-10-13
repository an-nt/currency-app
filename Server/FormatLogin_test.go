package Server

import (
	"io/ioutil"
	"io"
	"net/http"
	"testing"
)

type Mock struct{}

func (m *Mock) Login(user uint, pass string) (string, error) {
	return "abc", nil
}

func TestFormatLogin(t *testing.T) {
	var formatter LoginFormatter
	formatter.Handler = &Mock{}
	reqBody := `"ID"`
	ioutil.
	req := http.NewRequest("POST", "http://localhost:8080/v1/login", io.Reader)

}
