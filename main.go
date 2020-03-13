package main

import (
	"database/sql"
	"goerr/kerr"
	"goerr/usecase"
	"log"

	"github.com/cockroachdb/errors"
)

func main() {
	err := usecase.GetUser()
	handleError(err)

	log.Println("++++++++++++++++++++")

	err = usecase.GetAdmin()
	handleError(err)
}

func handleError(err error) {
	switch {
	case err == nil:
		log.Println("Yay! no error")
	case errors.Is(err, sql.ErrNoRows) && kerr.IsSecureKind(err):
		log.Println("PERMISSION DENIED")
		log.Printf("%+v\n", err)
	case errors.Is(err, sql.ErrNoRows):
		log.Println("User not found")
		log.Printf("%+v\n", err)
	default:
		log.Fatal(err)
	}
}
