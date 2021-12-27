package routes

import (
	"github.com/gorilla/mux"
	"weltraumschaf.de/battleship/internal/server/handler"
)

func NewConfiguredRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	registerArticleRoutes(r)
	registerGameRoutes(r)
	registerUserRoutes(r)

	return r
}

func registerArticleRoutes(r *mux.Router) {
	h := handler.NewArticleHandler()
	r.HandleFunc("/articles", h.AllArticles)
	r.HandleFunc("/article", h.CreateArticle).Methods("POST")
	r.HandleFunc("/article/{id}", h.SingleArticle)
	r.HandleFunc("/article/{id}", h.UpdateArticle).Methods("PUT")
	r.HandleFunc("/article/{id}", h.DeleteArticle).Methods("DELETE")
}

func registerGameRoutes(r *mux.Router) {
	h := handler.NewGameHandler()
	r.HandleFunc("/games", h.AllGames)
}

func registerUserRoutes(r *mux.Router) {
	h := handler.NewUserHandler()
	r.HandleFunc("/users", h.AllUsers)
}

