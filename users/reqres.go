package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
	}

	CreateUserResponse struct {
		Ok string `json:"ok"`
		Id string `json:"id"`
	}

	GetUserRequest struct {
		Id        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	GetUserResponse struct {
		Id        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"string"`
	}
)

func encodeResponse(
	ctx context.Context,
	rw http.ResponseWriter,
	response interface{},
) error {
	return json.NewEncoder(rw).Encode(response)
}

func decodeUserRequest(
	ctx context.Context,
	r *http.Request,
) (interface{}, error) {

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}

	return req, nil
}
