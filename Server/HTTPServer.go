package Server

import (
	"CurrencyApp/API"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	UserServer
	http.Server
}
type IHttpServer interface {
	StartServer()
}

func (sv *HttpServer) StartServer() {
	//define a.apiLogin
	if sv.IApi == nil {
		sv.IApi = &API.Api{}
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to CURRENCY EXCHANGE app"))
	})
	router.HandleFunc("/v1/login", sv.IApi.Login).Methods("POST")
	router.HandleFunc("/v1/exchangerates/{currencycode}", sv.IApi.GetExchangeRate).Methods("GET")

	sv.Handler = router
	sv.Addr = ":8080"
	sv.ReadTimeout = 15 * time.Second
	sv.WriteTimeout = 15 * time.Second
	log.Fatal(sv.ListenAndServe())
}
