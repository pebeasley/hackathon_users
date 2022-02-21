package users

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(
			ctx,
			req.Email,
			req.Password,
			req.Firstname,
			req.Lastname,
		)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("test from service")

		req := request.(GetUserRequest)
		user, err := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			Id:        fmt.Sprintf("%d", user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}, err
	}
}
