package Server

import (
	"CurrencyApp/API"
)

type UserServer struct {
	API.IApi
}

type IServer interface {
	StartServer()
}
