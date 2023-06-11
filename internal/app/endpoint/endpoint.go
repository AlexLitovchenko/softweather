package endpoint

import (
	"io/ioutil"
	"net/http"
)

type Service interface {
	LongestSubstring(string) string
}
type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}
func (e *Endpoint) Substring(w http.ResponseWriter, r *http.Request) {
	// Прочитаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusInternalServerError)
		return
	}

	// Преобразуем полученные данные в строку
	str := string(body)

	// Найдем самую длинную подстроку
	substring := e.s.LongestSubstring(str)

	// Отправим ответ клиенту
	w.Write([]byte(substring))
}
