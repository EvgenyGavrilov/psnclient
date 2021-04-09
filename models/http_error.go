package models

import "fmt"

// HTTPError структура ошибки от playstation store
type HTTPError struct {
	// StatusCode http статус ответа
	StatusCode int
	// Body тело ответа
	Body string
}


func (e *HTTPError) Error() string {
	return fmt.Sprintf("Response http status %d", e.StatusCode)
}
