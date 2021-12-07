package usecase

import (
	"github.com/taiki1288/Go-CleanArchitecture-APIServer/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

// User情報を作成して返すメソッド
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

// ユーザーの一致するidを探して返すメソッド
func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

// ユーザー情報を更新するメソッド
func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

// ユーザー情報を削除するメソッド
func (interactor *UserInteractor) DeleteByUser(user domain.User) (err error) {
	err = interactor.UserRepository.DeleteById(user)
	return
}