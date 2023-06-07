package api

import (
	"log"
	"net/http"
	"time"

	"githud.com/test-task/insert/kesh"
)

type Json struct {
	Domain string        `json:"domain"`
	Delay  time.Duration `json:"delay"`
}

type Request struct {
	Successful    uint `json:"successful"`
	NotSuccessful uint `json:"notSuccessful"`
}

type Pointer struct {
	Particula Request `json:"particula"`
	Min       Request `json:"min"`
	Max       Request `json:"max"`
}

type application struct {
	k    kesh.Kesh
	stat Pointer
}

func New(k kesh.Kesh) *application {
	return &application{
		k: k,
	}
}

// маршрутизатор
func (a *application) Routr() *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/site", validUrl(validMethd(a.particula)))
	mux.HandleFunc("/site/min", validUrl(validMethd(a.min)))
	mux.HandleFunc("/site/max", validUrl(validMethd(a.max)))
	mux.HandleFunc("/endpoint", validUrl(validMethd(identifie(a.pointer))))

	log.Println("server start")
	return mux
}
