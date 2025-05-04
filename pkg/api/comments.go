package api

import (
	"apiBlog/pkg/handlers"
	"net/http"
)

func commentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetComments(w, r)
	case http.MethodDelete:
		handlers.CreateComment(w, r)
	}
}