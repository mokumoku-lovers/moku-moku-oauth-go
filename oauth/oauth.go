package oauth

import (
	"net/http"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXPublic    = "X-Public"
	headerXCallerId  = "X-Caller-Id"
	paramAccessToken = "access_token"
)

var oauthRestClient = rest.RequestBuilder{
	BaseURL: "http://localhost:8080",
	Timeout: 200 * time.Millisecond,
}

type accessToken struct {
	Id     string `json:"id"`
	UserId int64  `json:"user_id"`
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}

	request.Header.Del(headerXCallerId)
}
