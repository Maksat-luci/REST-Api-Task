package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SubStr ...
type SubStr struct {
	Substring string `json:"substr"`
}
type findSubString interface {
	Find(str string) string
}

// Find метод для нахождения подстроки
func (s *SubStr) Find(str string) string {
	resultStr := ""
	for _, k := range str {
		cheker := true
		for i := 0; i < len(resultStr)-1; i++ {
			if resultStr[i] == byte(k) {
				cheker = false
				break
			}
		}
		if cheker {
			resultStr += string(k)
		}
	}

	return resultStr
}

// хэндлер типа middleWare
func (s *APIserver) handleSubString() http.HandlerFunc {
	var (
		err                  error
		intermediateVariable []byte
		objectSubstr         SubStr
	)

	return func(w http.ResponseWriter, r *http.Request) {
		// проверка получаемого метода
		if r.Method != "POST" {
			http.Error(w, "it`s not Post method", http.StatusMethodNotAllowed)
			return
		}
		// читаем с запроса его содержимое
		intermediateVariable, err = io.ReadAll(r.Body)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleSubString",
				"err":      err,
			}).Error()
			http.Error(w, "request not valid ", http.StatusBadRequest)
		}
		err = json.Unmarshal(intermediateVariable, &objectSubstr)

		if err != nil {
			// if error is not nil
			// print error
			http.Error(w, "bad Json", http.StatusBadRequest)
			return
		}
		// кастим массив байтов в стринг и присваеваем его структуре
		objectSubstr.Substring = objectSubstr.Find(objectSubstr.Substring)
		// отправляем ответ
		w.Write([]byte(objectSubstr.Substring))
		r.Body.Close()

	}
}
