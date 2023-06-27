package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *application) certain(w http.ResponseWriter, r *http.Request) {
	result, err := a.k.Single(r.FormValue("domain"))
	if err != nil {
		log.Println("handler certain:\t", err)
		http.Error(w, "not found", http.StatusBadRequest)

		a.stat.Certain.NotSuccessful++
		return
	}

	j, err := json.Marshal(&Json{Domain: result.Domain, Delay: result.Delay})
	if err != nil {
		log.Println("handler certain:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Certain.NotSuccessful++
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	_, err = w.Write(j)
	if err != nil {
		log.Println("handler certain:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Certain.NotSuccessful++
		return
	}

	w.WriteHeader(http.StatusOK)

	a.stat.Certain.Successful++
}

func (a *application) min(w http.ResponseWriter, r *http.Request) {
	result, err := a.k.Min()
	if err != nil {
		log.Println("handler min:\t", err)
		http.Error(w, "", http.StatusInternalServerError)

		a.stat.Min.NotSuccessful++
		return
	}

	j, err := json.Marshal(&Json{Domain: result.Domain, Delay: result.Delay})
	if err != nil {
		log.Println("handler min:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Min.NotSuccessful++
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	_, err = w.Write(j)
	if err != nil {
		log.Println("handler min:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Min.NotSuccessful++
		return
	}

	w.WriteHeader(http.StatusOK)

	a.stat.Min.Successful++
}

func (a *application) max(w http.ResponseWriter, r *http.Request) {
	result, err := a.k.Max()
	if err != nil {
		log.Println("handler max:\t", err)
		http.Error(w, "", http.StatusInternalServerError)

		a.stat.Max.NotSuccessful++
		return
	}

	j, err := json.Marshal(&Json{Domain: result.Domain, Delay: result.Delay})
	if err != nil {
		log.Println("handler max:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Max.NotSuccessful++
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	_, err = w.Write(j)
	if err != nil {
		log.Println("handler max:\t", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		a.stat.Max.NotSuccessful++
		return
	}

	w.WriteHeader(http.StatusOK)

	a.stat.Max.Successful++
}

func (a *application) pointer(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(a.stat)
	if err != nil {
		log.Println("handler max:\t", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	_, err = w.Write(j)
	if err != nil {
		log.Println("handler max:\t", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
