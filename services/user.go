package services

import (
	"context"

	"github.com/rickybell/go-interfaces/app/entities"
	"github.com/rickybell/go-interfaces/app/interfaces"
	"github.com/rickybell/go-interfaces/app/repositories"
)

// UserServiceInterface is an interface for the UserService
type ServiceInterface interface {
	All() ([]interfaces.Model, error)
	// Find(id int) (*interfaces.User, error)
	Create(model *interfaces.Model) (*interfaces.Model, error)
	// Update(user *interfaces.User) (*interfaces.User, error)
	// Delete(id int) error
}

// UserService is a struct that represents the UserService
type UserService struct {
	Repository repositories.UserPostgresSqlGormRepository
	ctx        context.Context
}

// NewUserService is a function that returns a new UserService
func NewUserService(userRepository repositories.UserPostgresSqlGormRepository) *UserService {
	return &UserService{
		Repository: userRepository,
	}
}

func (u *UserService) All() ([]entities.User, error) {
	users, err := u.Repository.All(u.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) Create(user entities.User) (*entities.User, error) {
	users, err := u.Repository.Create(u.ctx, user)
	if err != nil {
		return nil, err
	}

	return users, nil
}
