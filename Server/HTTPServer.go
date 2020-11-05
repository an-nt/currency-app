package Server

import (
	"CurrencyApp/API"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	http.Server
	Exec API.IApi
}
type IHttpServer interface {
	StartServer()
}

func (sv *HttpServer) StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to CURRENCY EXCHANGE app"))
	})
	router.HandleFunc("/v1/account", sv.formatSignUp).Methods("POST")

	router.HandleFunc("/v1/login", sv.formatLogin).Methods("POST")

	router.HandleFunc("/v1/exchangerates/{currencycode}", sv.formatGetExRate).Methods("GET")

	sv.Handler = router
	sv.Addr = ":8080"
	sv.ReadTimeout = 15 * time.Second
	sv.WriteTimeout = 15 * time.Second
	log.Fatal(sv.ListenAndServe())
}
