package controllers

import (
	"fmt"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/domain/items"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/services"
	"net/http"
)

var(
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, request *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (controller *itemsController) Create(w http.ResponseWriter, request *http.Request) {
	if err := oath.AuthenticateRequest(request); err != nil {
		return
	}

	item := items.Item{
		Seller: oath.GetCallerId(request),
	}

	result, err := services.ItemsService.Create(item)

	if err != nil {
		return
	}
	fmt.Println(result)
}

func (controller *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hi"))
}
