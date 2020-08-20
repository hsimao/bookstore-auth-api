package db

import (
	"github.com/hsimao/bookstore_auth-api/src/domain/access_token"
	"github.com/hsimao/bookstore_users-api/utils/errors"
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
	return nil, nil
}
