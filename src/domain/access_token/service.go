package access_token

import "github.com/hsimao/bookstore_users-api/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string) (AccessToken, *errors.RestErr) {
	return nil, nil
}
