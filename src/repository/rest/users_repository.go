package rest

import (
	"UdemyApp/bookstore_oauth-api/src/domain/users"
	"UdemyApp/bookstore_oauth-api/src/utils/errors"
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	_ "github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "http://127.0.0.1:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when tryit to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user reponse")
	}

	return &user, nil
}