package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetCountriesQuery(db *sql.DB) []Country {
	res := Country{}
	rows, err := db.Query("SELECT * FROM countries")
	CheckError(err)
	defer rows.Close()

	var (
		name     string
		language string
	)

	for rows.Next() {
		rows.Scan(&name, &language)
		res.Language = language
		res.Name = name
		countries = append(countries, res)
	}
	fmt.Println(countries)
	return countries
}

func AddCountryQuery(country *Country, db *sql.DB) (string, error) {
	var err error
	var msg string
	if country != nil {
		_, err = db.Exec("INSERT INTO countries VALUES ($1,$2)", country.Name, country.Language) // probar luego con :1 y :2
		if err != nil {
			msg = "Error ejecutando la Query: "
			return msg, err
		}
		msg = "Registro guardado con exito "
		fmt.Printf("Registro con exito %v", country)
	}
	return msg, err
}
