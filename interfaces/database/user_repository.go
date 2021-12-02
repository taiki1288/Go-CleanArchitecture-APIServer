package database

import "Go-CleanArchitecture-APIServer/domain"

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = db.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (db *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = db.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (db *UserRepository) FindAll() (users domain.Users, err error) {
	if err = db.Find(&users).Error; err != nil {
		return
	}
	return
}