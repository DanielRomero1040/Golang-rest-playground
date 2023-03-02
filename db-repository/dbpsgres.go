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

type TokenToStorage struct {
	Token    string
	UserName string
	Date     string
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

func DbClose() {
	err := Dbp.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}

}
