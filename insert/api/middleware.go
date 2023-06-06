package api

import (
	"net/http"
	"regexp"

	confi "git_p/githud.com/test-task/insert"
)

var validPath = regexp.MustCompile("^/(site|endpoint)?(/max|/min)?$")

// роверка адреса запроса
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

// индификация админа
func identifie(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Form.Get("key")

		if key != confi.Key {
			http.Error(w, "denial of access", http.StatusBadRequest)
		}

		f(w, r)
	}
}
