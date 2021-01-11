package rest

import (
	"flag"
	"github.com/mercadolibre/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	//rest.FlushMockups()
	//rest.AddMockups(&rest.Mock{
	//	URL: "http://127.0.0.1:8080/users/login",
	//	HTTPMethod: http.MethodPost,
	//	ReqBody: `{"email": "xxx@gmail.com", "password": "password"}`,
	//	RespHTTPCode: -1,
	//	RespBody: `{}`,
	//})
	//repository := usersRepository{}
	//user, err := repository.LoginUser("xxx@gmail.com", "password")
	//
	//assert.Nil(t, user)
	//assert.NotNil(t, err)
	//assert.EqualValues(t, http.StatusInternalServerError, err.Status)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}



