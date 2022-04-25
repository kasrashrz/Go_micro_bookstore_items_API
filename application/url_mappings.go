package application

import (
	"github.com/kasrashrz/Go_micro_bookstore_items_API/controllers"
	"net/http"
)

func MapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.ItemsController.Get).Methods(http.MethodGet)
}
