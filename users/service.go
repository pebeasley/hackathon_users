package users

import (
	"context"
	"log"

	"github.com/pebeasley/users/database"
)

type Service interface {
	CreateUser(
		ctx context.Context,
		firstName string,
		lastName string,
		email string,
		password string,
	) (string, error)

	GetUser(ctx context.Context, id string) (User, error)
}

type service struct {
	logger log.Logger
}

func NewService() Service {
	return &service{
		logger: *log.Default(),
	}
}

func (s service) CreateUser(
	ctx context.Context,
	firstName string,
	lastName string,
	email string,
	password string,
) (string, error) {
	err := database.DB.Create(User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return "nil", err.Error
	}
	return email, nil
}

func (s service) GetUser(ctx context.Context, id string) (User, error) {
	user := &User{}
	result := database.DB.First(user, id)
	return *user, result.Error
}
