package endpoint

import (
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error with reading request", http.StatusInternalServerError)
		return
	}

	str := string(body)

	substring := e.s.LongestSubstring(str)

	w.Write([]byte(substring))
}
