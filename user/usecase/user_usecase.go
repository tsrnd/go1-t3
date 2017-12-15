package usecase

import (
	model "github.com/goweb3/user"
	repos "github.com/goweb3/user/repository"
)

// UserUsecase interface
type UserUsecase interface {
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetPrivateUserDetailsByEmail(email string) (*model.PrivateUserDetails, error)
	Create(email, name, password string) (int64, error)
}

type userUsecase struct {
	userRepos repos.UserRepository
}

func (a *userUsecase) GetByID(id int64) (*model.User, error) {
	return a.userRepos.GetByID(id)
}

func (a *userUsecase) GetByEmail(email string) (*model.User, error) {
	return a.userRepos.GetByEmail(email)
}

func (a *userUsecase) GetPrivateUserDetailsByEmail(email string) (*model.PrivateUserDetails, error) {
	return a.userRepos.GetPrivateUserDetailsByEmail(email)
}

func (a *userUsecase) Create(email, name, password string) (int64, error) {
	return a.Create(email, name, password)
}

// NewUserUsecase func
func NewUserUsecase(a repos.UserRepository) UserUsecase {
	return &userUsecase{a}
}
