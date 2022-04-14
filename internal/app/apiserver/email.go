package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// EmailTask ...
type EmailTask struct {
	Emailstr     string `json:"email"`
	FindedEmails string
}

// EmailFind ...
type EmailFind interface {
	FindEmail(email string) string
}

// SortSpaceAndSleshN функция для обработки стринг
func SortSpaceAndSleshN(str string) string {
	variable := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			continue
		} else if str[i] == ' ' {
			continue
		} else {
			variable += string(str[i])
		}
	}
	return variable

}

// FindEmail ...
func (m *EmailTask) FindEmail(str string) string {
	intermediateVariable := SortSpaceAndSleshN(str)
	result := strings.Split(intermediateVariable, "Email:")[1]
	return result

}

func (s *APIserver) handleEmail() http.HandlerFunc {
	var (
		ObjectEmail          EmailTask
		intermediateVariable []byte
		err                  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}
		intermediateVariable, err = io.ReadAll(r.Body)
		fmt.Println(string(intermediateVariable))
		if err != nil {
			s.logger.Error("Bad request from r.body ")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		err = json.Unmarshal(intermediateVariable, &ObjectEmail)
		if err != nil {
			s.logger.Error("not valid json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ObjectEmail.FindedEmails = ObjectEmail.FindEmail(ObjectEmail.Emailstr)
		w.Write([]byte(ObjectEmail.FindedEmails))

	}
}
