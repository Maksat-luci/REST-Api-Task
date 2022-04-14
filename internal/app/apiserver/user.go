package apiserver

import (
	"encoding/json"
	"io"
	"net/http"
)

// User структура юзер для 4го таска
type User struct {
	UserID    int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *APIserver) handleCreateUser() http.HandlerFunc {
	var (
		intermediateVariable []byte
		err                  error
		ObjectUser           User
	)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method POST not allowed", http.StatusMethodNotAllowed)
			return
		}
		intermediateVariable, err = io.ReadAll(r.Body)
		err = json.Unmarshal(intermediateVariable, &ObjectUser)
		if err != nil {
			http.Error(w, "invalid json format", http.StatusBadRequest)
			return
		}

	}
}
