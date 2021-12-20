package main

import (
	"fmt"
	"os"
	"weltraumschaf.de/battleship/internal/server"
)

func main() {
	var app = server.Create()
	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
