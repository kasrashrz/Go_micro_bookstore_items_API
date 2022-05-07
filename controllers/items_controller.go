package controllers

import (
	"encoding/json"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath/errors"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/domain/items"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/services"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/utils/http_utils"
	"io/ioutil"
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
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		respErr := errors.BadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}

	defer request.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.BadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oath.GetCallerId(request)

	result, createErr := services.ItemsService.Create(itemRequest)

	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (controller *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
