package controller

import (
	"api-rest/service"
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !isAllowedMethod(w, r, http.MethodGet) {
			return
		}
		service.Index(w, r)
	})

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			service.GetCountries(w, r)
		case http.MethodPost:
			service.AddCountry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if !isAllowedMethod(w, r, http.MethodPost) {
			return
		}
		service.GetToken(w, r)
	})
}

func isAllowedMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	response := true
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		response = false
	}
	return response
}
