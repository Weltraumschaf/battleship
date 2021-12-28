package router

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

type httpMethod string

const (
	httpGet httpMethod = "GET"
	httpPost = "POST"
	httpPut = "PUT"
	httpDelete = "DELETE"
)

func registerArticleRoutes(r *mux.Router) {
	h := handler.NewArticleHandler()
	r.HandleFunc("/articles", h.AllArticles).Methods(string(httpGet))
	r.HandleFunc("/article", h.CreateArticle).Methods(string(httpPost))
	r.HandleFunc("/article/{id}", h.SingleArticle).Methods(string(httpGet))
	r.HandleFunc("/article/{id}", h.UpdateArticle).Methods(string(httpPut))
	r.HandleFunc("/article/{id}", h.DeleteArticle).Methods(string(httpDelete))
}

func registerGameRoutes(r *mux.Router) {
	h := handler.NewGameHandler()
	r.HandleFunc("/games", h.AllGames).Methods(string(httpGet))
}

func registerUserRoutes(r *mux.Router) {
	h := handler.NewUserHandler()
	r.HandleFunc("/users", h.AllUsers).Methods(string(httpGet))
	r.HandleFunc("/user", h.CreateUsers).Methods(string(httpPost))
	r.HandleFunc("/user/{name}", h.SingleUser).Methods(string(httpGet))
	r.HandleFunc("/user/{name}", h.UpdateUser).Methods(string(httpPut))
	r.HandleFunc("/user/{name}", h.DeleteUser).Methods(string(httpDelete))
}

