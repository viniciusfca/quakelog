package errors

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

// error.go
func SendErrorResponse(w http.ResponseWriter, err *APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	body, _ := json.Marshal(map[string]string{
		"message": err.Message,
	})
	w.Write(body)
}
