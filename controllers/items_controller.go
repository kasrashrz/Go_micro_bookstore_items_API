package controllers

import (
	"fmt"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/domain/items"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/services"
	"net/http"
)

func Create(w http.ResponseWriter, request *http.Request) {
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

func Get(w http.ResponseWriter, r *http.Request) {

}
