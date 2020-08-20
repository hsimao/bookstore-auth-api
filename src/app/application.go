package app

import(
	"github.com/hsimao/bookstore_auth-api/src/domain/access_token"
	"github.com/hsimao/bookstore_auth-api/src/repository/db"
	)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
}
