package api

import (
	"net/http"
	"time"
)

type json struct {
	url   string        `json:"url"`
	delay time.Duration `json:"delay"`
}

func particula(w http.ResponseWriter, r *http.Request) {

}

func min(w http.ResponseWriter, r *http.Request) {

}

func max(w http.ResponseWriter, r *http.Request) {

}

func pointer(w http.ResponseWriter, r *http.Request) {

}
