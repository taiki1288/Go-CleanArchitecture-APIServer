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