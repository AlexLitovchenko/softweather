package main

import (
	"log"

	"github.com/AlexLitovchenko/softweather/task/internal/app/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	a.Run()
}
