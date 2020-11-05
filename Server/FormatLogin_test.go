package Server

import (
	"testing"
)

type Mock struct{}

func (m *Mock) Login(user uint, pass string) (string, error) {
	return "abc", nil
}

func TestFormatLogin(t *testing.T) {

}
