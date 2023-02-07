package main

import (
	vault "api-rest/resources"
	"api-rest/server"
)

func main() {
	vault.VaultConfig()
	srv := server.New(":8080")
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
