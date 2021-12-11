package battleship

import (
    "fmt"
    "github.com/urfave/cli/v2"
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
    fmt.Println(NewBoard())
    return nil
}
