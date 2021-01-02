package app

import (
	"UdemyApp/bookstore_oauth-api/src/client/cassandra"
	"UdemyApp/bookstore_oauth-api/src/domain/access_token"
	"UdemyApp/bookstore_oauth-api/src/http"
	"UdemyApp/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHanlder := http.NewHanlder(atService)

	router.GET("/oauth/access_token/:access_token_id", atHanlder.GetById)

	router.Run(":8080")
}