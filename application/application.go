package application

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	MapUrls()
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout: 2 * time.Second,
		IdleTimeout: 60 * time.Second,
		Addr:        "127.0.0.1:8080",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
