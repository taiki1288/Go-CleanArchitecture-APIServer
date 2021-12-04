package controllers

import "Go-CleanArchitecture-APIServer/usecase"

type UserController struct {
	Interactor usecase.UserInteractor
}