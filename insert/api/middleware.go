package api

import (
	"net/http"
	"regexp"

	confi "githud.com/test-task/insert"
)

var validPath = regexp.MustCompile("^/(site|endpoint)?(/max|/min)?$")

// поверка адреса запроса
func validUrl(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Error(w, "invalid path", http.StatusBadRequest)
			return
		}

		f(w, r)
	}
}

// проверка метода
func validMethd(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		f(w, r)
	}
}

// индификация админа
func identifie(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if key == confi.Key {
			f(w, r)
		} else {
			http.Error(w, "denial of access", http.StatusBadRequest)
			return
		}
	}
}
