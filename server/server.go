package server

import (
	"api-rest/controller"
	"net/http"
)

func New(addr string) *http.Server {
	controller.InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
