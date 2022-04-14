package apiserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleSubString(t *testing.T) {
	s := New(NewConfig())
	// newConfig возвращает конфиг для New  тот в свою очередь возвращает apiServer
	rec := httptest.NewRecorder()
	// создаем RESPONSE
	strJSON := []byte(`{"substr":"abca"}`)
	req, _ := http.NewRequest(http.MethodPost, "/rest/substr/find", bytes.NewBuffer(strJSON))
	// создаём реквест (запрос)
	s.handleSubString().ServeHTTP(rec, req)

	// отправляем методу handleHello наш запрос с ответом(пустым)
	assert.Equal(t, "abc", rec.Body.String())
	
	// исправитть тест не находит ответ от хендлера
	// сравниваем ответ с помощью библеотеки тестифай, ответ от нашего запроса и зравниваем его с ожидаемым результатом
}
