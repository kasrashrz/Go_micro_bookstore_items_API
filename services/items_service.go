package services

import (
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath/errors"
	"github.com/kasrashrz/Go_micro_bookstore_items_API/domain/items"
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct{}

func (service itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, nil
}

func (service itemsService) Get(itemName string) (*items.Item, **errors.RestErr) {
	return nil, nil
}
