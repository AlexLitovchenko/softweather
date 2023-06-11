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
	// Зарегистрируем обработчик для endpoint'а /api/substring
	http.HandleFunc("/api/substring", a.e.Substring)
	return a, nil
}

func (a *App) Run() {
	// Запустим HTTP-сервер на порту 8080
	log.Println("HTTP-сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
