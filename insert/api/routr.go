package api

import "net/http"

// маршрутизатор
func Routr() *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/site", validUrl(particula))
	mux.HandleFunc("/site/min", validUrl(min))
	mux.HandleFunc("/site/max", validUrl(max))
	mux.HandleFunc("/pointer", validUrl(identifie(pointer)))

	return mux
}
