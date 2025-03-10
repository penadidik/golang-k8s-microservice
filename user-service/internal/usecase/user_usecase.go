package usecase

import (
	"errors"
	"user-service/internal/entity"
	"user-service/internal/repository"
)

type UserUseCase interface {
	RegisterUser(user *entity.User) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo}
}

func (uc *userUseCase) RegisterUser(user *entity.User) error {
	existingUser, _ := uc.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}
	return uc.repo.Create(user)
}

