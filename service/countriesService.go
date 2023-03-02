package service

import (
	db "api-rest/db-repository"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	db.DoPostgress()
	countries2 := db.GetCountriesQuery(db.Dbp)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(countries2)
}

func AddCountry(w http.ResponseWriter, r *http.Request) {
	db.DoPostgress()
	country := &db.Country{} // revisar
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	msg, err := db.AddCountryQuery(country, db.Dbp)
	// countries = append(countries, country)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		json.NewEncoder(w).Encode(country)
		fmt.Fprint(w, msg)
	}
}

func SessionMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//validación de session
		var store = sessions.NewCookieStore([]byte(os.Getenv("")))
		session, _ := store.Get(r, "session-user")
		if session.IsNew {
			fmt.Fprintf(w, "Debes iniciar sesión para poder ver está página")
			return
		}
		if !isTokenSessionValid(session) {
			fmt.Printf("session.Values: %v\n", session.Values)
			return
		}
		fmt.Printf("session.Values: %v\n", session.Values)
		f(w, r)
	}
}
