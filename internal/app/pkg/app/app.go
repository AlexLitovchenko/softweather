package app

import (
	"log"
	"net/http"

	"github.com/AlexLitovchenko/softweather/task/internal/app/endpoint"
	"github.com/AlexLitovchenko/softweather/task/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	http.HandleFunc("/api/substring", a.e.Substring)
	return a, nil
}

func (a *App) Run() {
	log.Println("HTTP-server started on port - 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
