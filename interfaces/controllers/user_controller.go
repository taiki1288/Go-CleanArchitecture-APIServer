package controllers

import (
	"Go-CleanArchitecture-APIServer/domain"
	"Go-CleanArchitecture-APIServer/interfaces/database"
	"Go-CleanArchitecture-APIServer/usecase"

	"strconv"
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

// endpoint POST /users
func (controller *UserController) CreateUser(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// endpoint GET /users/:id/
func (controller *UserController) GetUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// endpoint GET /users/
func (controller *UserController) GetUsers(c Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// endpoint UPDATE /users/:id/
func (controller *UserController) UpdateUser(c Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをUser構造体のIDフィールドに格納
	user := domain.User{ID: id}
	// userをUpdate()に代入
	user, err = controller.Interactor.Update(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// endpoint DELETE /users/:id/
func (controller *UserController) DeleteUser(c Context) (err error) {
	// Atoiでc.Param("id")をint型のidに変換
	id, _ := strconv.Atoi(c.Param("id"))
	// idをUser構造体のIDフィールドに格納
	user := domain.User{ID: id}
	// userをDeleteByUser()に代入
	err = controller.Interactor.DeleteByUser(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}