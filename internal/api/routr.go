package api

import (
	"log"
	"net/http"
	"time"

	"githud.com/test-task/internal/cache"
)

// структур адля ответа клиеннту
type Json struct {
	Domain string        `json:"domain"`
	Delay  time.Duration `json:"delay"`
}

// структура хронит в себе количество обращения к эндпоиту
type Request struct {
	Successful    uint `json:"successful"`
	NotSuccessful uint `json:"notSuccessful"`
}

// структура хронит все обращения к эндпоиторам
type Pointer struct {
	Certain Request `json:"certain"`
	Min     Request `json:"min"`
	Max     Request `json:"max"`
}

type application struct {
	k    cache.Cache
	stat Pointer
}

func New(k cache.Cache) *application {
	return &application{
		k: k,
	}
}

// маршрутизатор
func (a *application) Router() *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/site", validUrl(validMethod(a.certain)))
	mux.HandleFunc("/site/min", validUrl(validMethod(a.min)))
	mux.HandleFunc("/site/max", validUrl(validMethod(a.max)))
	mux.HandleFunc("/endpoint", validUrl(validMethod(identified(a.pointer))))

	log.Println("server start")
	return mux
}
