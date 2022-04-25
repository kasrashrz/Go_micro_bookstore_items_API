package controllers

type usersControllerInterface interface {
	CreateUser()
}

type usersController struct{}

var UsersController usersControllerInterface = &usersController{}

func (controller *usersController) CreateUser() {

}
