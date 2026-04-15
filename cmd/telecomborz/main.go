package main

import (
    "log"

    "github.com/yourname/telecomborz/internal/app"
)

func main() {
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
