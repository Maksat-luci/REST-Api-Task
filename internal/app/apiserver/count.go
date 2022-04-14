package apiserver

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *APIserver) handleCountSub() http.HandlerFunc {
	var (
		counter int
		sub     int
		err     error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "NOT valid method", http.StatusMethodNotAllowed)
			return
		}
		sub, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[4])
		counter, err = s.red.Get("counter").Int()
		if err != nil {
			http.Error(w, "Sorry something is wrong with the server", http.StatusInternalServerError)
			return
		}
		s.red.Set("counter", counter-sub, 0)
		w.Write([]byte("OK reduced successfully"))
	}
}

func (s *APIserver) handleCounterValue() http.HandlerFunc {

	var (
		val    int
		err    error
		result string
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Not Valid method", http.StatusMethodNotAllowed)
			return
		}
		val, err = s.red.Get("counter").Int()
		result = strconv.Itoa(val)
		if err != nil {
			http.Error(w, "Sorry something is wrong with the server", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
	}
}

func (s *APIserver) handleCountAdd() http.HandlerFunc {
	var (
		counter int
		add     int
		err     error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "bratik yzai POSt Method", http.StatusMethodNotAllowed)
			return
		}
		add, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[4])
		counter, err = s.red.Get("counter").Int()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.red.Set("counter", counter+add, 0)
		w.Write([]byte("OK enlarged successfully"))
	}
}


// написать юнит тесты
// закончить 4тый таск
