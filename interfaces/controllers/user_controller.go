package controllers

import (
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