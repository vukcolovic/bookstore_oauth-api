package db

import (
	"UdemyApp/bookstore_oauth-api/src/client/cassandra"
	"UdemyApp/bookstore_oauth-api/src/domain/access_token"
	"UdemyApp/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("database connection error")
}


