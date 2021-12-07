package usecase

import "github.com/taiki1288/Go-CleanArchitecture-APIServer/domain"

type UserRepository interface {
	Store(domain.User) (domain.User, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}
