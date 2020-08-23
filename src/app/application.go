package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hsimao/bookstore_auth-api/src/domain/access_token"
	"github.com/hsimao/bookstore_auth-api/src/http"
	"github.com/hsimao/bookstore_auth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
