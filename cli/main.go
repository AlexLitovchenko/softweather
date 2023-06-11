package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Получаем аргументы командной строки
	if len(os.Args) < 3 {
		log.Fatal("Необходимо указать строку и URL в качестве аргументов")
	}

	str := os.Args[1]
	url := os.Args[2]

	// Отправляем запрос на указанный URL с переданной строкой
	resp, err := http.Post(url, "text/plain", strings.NewReader(str))
	if err != nil {
		log.Fatal("Не удалось отправить запрос:", err)
	}
	defer resp.Body.Close()

	// Читаем ответ сервера
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Не удалось прочитать ответ сервера:", err)
	}

	// Выводим ответ на экран
	fmt.Println(string(body))
}
