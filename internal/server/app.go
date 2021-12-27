package server

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"net/http"
	"weltraumschaf.de/battleship/internal/server/handler"
)

func Create() *cli.App {
	return &cli.App{
		Name:    "battleship",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Sven Strittmatter",
				Email: "ich@weltraumschaf.de",
			},
		},
		Action: Execute,
	}
}

func Execute(c *cli.Context) error {
	r := mux.NewRouter().StrictSlash(true)

	registerArticleRoutes(r)
	registerGameRoutes(r)
	registerUserRoutes(r)

	return http.ListenAndServe(":10000", r)
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

