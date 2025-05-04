package api

import (
	"net/http"

	"apiBlog/pkg/handlers"
)

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetPosts(w, r)
	case http.MethodPost:
		handlers.CreatePost(w, r)
	}
}
