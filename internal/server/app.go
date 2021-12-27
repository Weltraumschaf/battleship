package server

import (
	"github.com/urfave/cli/v2"
	"net/http"
	"weltraumschaf.de/battleship/internal/server/routes"
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
	r := routes.NewConfiguredRouter()

	return http.ListenAndServe(":10000", r)
}
