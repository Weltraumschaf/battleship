package server

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"log"
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
	log.Fatal(http.ListenAndServe(":10000", r))

	return nil
}

func registerArticleRoutes(r *mux.Router) {
	h := handler.NewArticleHandler()
	r.HandleFunc("/", h.HomePage)
	r.HandleFunc("/articles", h.ReturnAllArticles)
	r.HandleFunc("/article", h.CreateNewArticle).Methods("POST")
	r.HandleFunc("/article/{id}", h.ReturnSingleArticle)
	r.HandleFunc("/article/{id}", h.DeleteArticle).Methods("DELETE")
}
