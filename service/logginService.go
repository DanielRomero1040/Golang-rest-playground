package service

import (
	db "api-rest/db-repository"
	vault "api-rest/resources"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gorilla/sessions"
)

// Loggin - TODO - cambiar nombre del metodo
func GetToken(w http.ResponseWriter, r *http.Request) {
	db.DoPostgress()

	user := &db.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	if !db.ValidateUser(user, db.Dbp) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Usuario y/o Contraseña Invalidos")
		return
	}

	date := time.Now().Format(time.RFC3339Nano)
	signedToken := generateSignedToken(user, date)

	tokenToStorage := &db.TokenToStorage{
		Token:    signedToken,
		UserName: user.UserName,
		Date:     date,
	}

	db.SaveTokenQuery(tokenToStorage, db.Dbp)

	var store = sessions.NewCookieStore([]byte(os.Getenv("PROJECT_ENV")))
	store.MaxAge(300)
	session, _ := store.Get(r, "session-user")
	session.Values["jwt"] = signedToken
	session.Values["username"] = user.UserName
	session.Options.MaxAge = 300
	erro := session.Save(r, w)
	if erro != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, signedToken)
}

func generateSignedToken(user *db.User, date string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     user.UserName,
		"password": user.Password,
		"date":     date,
	})
	fmt.Println("token prueba", token)
	secret := []byte(vault.JwtSecret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("error firmando el token")
	}

	return signedToken
}

func isTokenSessionValid(session *sessions.Session) bool {
	jwtReceived := session.Values["jwt"].(string)
	token, err := jwt.Parse(jwtReceived, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		mySecret := []byte(vault.JwtSecret)
		return mySecret, nil
	})
	//TODO key is of invalid type ... debug el metodo parse! *** solución, colocar el string tipo []byte como dice la documentacion
	fmt.Println(token, err)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"], claims["password"], claims["date"])
	} else {
		fmt.Println(err)
	}

	return token.Valid
}
