package api

import (
	"github.com/go-chi/chi"
)

var Rout = chi.NewRouter()

func Init() {
	// Роуты для постов
	Rout.HandleFunc("/posts", postsHandler)
	Rout.HandleFunc("/post/{id}", postHandler)

	// Роут для комментов
	Rout.HandleFunc("/posts/{id}/comments", commentHandler)
}
