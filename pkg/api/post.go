package api

import (
	"net/http"

	"apiBlog/pkg/handlers"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetPost(w, r)
	case http.MethodDelete:
		handlers.DeletePost(w, r)
	}
}
