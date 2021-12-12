package main

import (
	"log"

	"github.com/ayupov-ayaz/todo/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
