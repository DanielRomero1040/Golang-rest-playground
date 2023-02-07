package db

import (
	props "api-rest/resources"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Country struct {
	Name     string
	Language string
}

type User struct {
	UserName string
	Password string
}

var countries []Country = []Country{}

var Dbp *sql.DB

func DoPostgress() {

	var err error
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", props.DbHost, props.DbPort, props.DbUser, props.DbPassword, props.DbName)

	// open database
	Dbp, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = Dbp.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// func handleError(msg string, err error) {
// 	if err != nil {
// 		fmt.Println(msg, err)
// 		panic(err)
// 	}
// }

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

func SaveTokenQuery(token string, user string, date string, db *sql.DB) (string, error) {
	var err error
	var msg string
	if token != "" {
		_, err = db.Exec("INSERT INTO tokens (TOKEN,USERLOG,DATE) VALUES ($1,$2,$3)", token, user, date) // probar luego con :1 y :2
		if err != nil {
			msg = "Error ejecutando la Query: "
			return msg, err
		}
		msg = "Token guardado con exito "
		fmt.Printf("Token guardado con exito %v", token)
	}
	return msg, err
}

func ValidateUser(user *User, db *sql.DB) bool {
	userDB := User{}
	if user != nil {
		rows, err := db.Query("SELECT username,password FROM users WHERE username = $1 ", user.UserName)
		CheckError(err)
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&userDB.UserName, &userDB.Password)
		}
		if *user == userDB {
			return true
		}
	}
	return false
}

func DbClose() {
	err := Dbp.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}

}
