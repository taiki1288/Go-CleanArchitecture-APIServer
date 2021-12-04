package controllers

import (
	"Go-CleanArchitecture-APIServer/domain"
	"Go-CleanArchitecture-APIServer/interfaces/database"
	"Go-CleanArchitecture-APIServer/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlhandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlhandler,
			},
		},
	}
}

func (controller *UserController) CreateUser(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
	}
	c.JSON(201, user)
	return
}