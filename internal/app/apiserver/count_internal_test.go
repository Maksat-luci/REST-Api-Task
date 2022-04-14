package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_handleCountAdd(t *testing.T) {
	s := New(NewConfig())
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/rest/counter/add/12", nil)
	s.handleCountAdd().ServeHTTP(res, req)
	assert.Equal(t, "OK enlarged successfully", res.Body.String())

}

func TestApiServer_handleCountSub(t *testing.T) {
	s := New(NewConfig())
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/rest/counter/sub/12", nil)
	s.handleCountSub().ServeHTTP(res,req)
	assert.Equal(t, "OK reduced successfully",res.Body.String())
}

func TestApiServer_handleCountVal(t *testing.T) {
	s := New(NewConfig())
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/rest/counter/val/",nil)
	s.handleCounterValue().ServeHTTP(res,req)
	assert.Equal(t, "0", res.Body.String())
}
