package main

import (
    "fmt"
    "os"
    "weltraumschaf.de/battleship/internal/battleship"
)

func main() {
    var app = battleship.Create()
    err := app.Run(os.Args)

    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }
}
