package app

import (
	"UdemyApp/bookstore_oauth-api/src/client/cassandra"
	"UdemyApp/bookstore_oauth-api/src/http"
	"UdemyApp/bookstore_oauth-api/src/repository/db"
	"UdemyApp/bookstore_oauth-api/src/repository/rest"
	"UdemyApp/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session:= cassandra.GetSession()
	session.Close()

	atService := access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository())
	atHanlder := http.NewHanlder(atService)

	router.GET("/oauth/access_token/:access_token_id", atHanlder.GetById)
	router.POST("/oauth/access_token", atHanlder.Create)

	router.Run(":8080")
}