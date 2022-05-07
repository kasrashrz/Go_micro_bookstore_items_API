package controllers

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath/errors"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/domain/items"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/services"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/utils/http_utils"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, request *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (controller *itemsController) Create(w http.ResponseWriter, request *http.Request) {
	if err := oath.AuthenticateRequest(request); err != nil {
		newErr := errors.RestErr{
			Message: err.Message(),
			Status:  err.Status(),
			Error:   err.Error(),
		}
		http_utils.RespondError(w, &newErr)
		return
	}

	item := items.Item{
		Seller: oath.GetCallerId(request),
	}

	result, err := services.ItemsService.Create(item)

	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (controller *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
