package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SaveTokenQuery(tokenToStorage *TokenToStorage, db *sql.DB) (string, error) {
	var err error
	var msg string
	if tokenToStorage.Token != "" {
		_, err = db.Exec("INSERT INTO tokens (TOKEN,USERLOG,DATE) VALUES ($1,$2,$3)", tokenToStorage.Token, tokenToStorage.UserName, tokenToStorage.Date) // probar luego con :1 y :2
		if err != nil {
			msg = "Error ejecutando la Query: "
			return msg, err
		}
		msg = "Token guardado con exito "
		fmt.Printf("Token guardado con exito %v", tokenToStorage.Token)
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
