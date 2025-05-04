package server

import (
	"net/http"
	"fmt"

	"apiBlog/pkg/api"
)

func Run(port int) error {
	api.Init()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), api.Rout)
}