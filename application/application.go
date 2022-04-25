package application

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	MapUrls()
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
