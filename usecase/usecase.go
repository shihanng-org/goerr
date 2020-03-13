package usecase

import (
	"goerr/db"
	"goerr/kerr"
)

func GetUser() error {
	return db.GetUser()
}

func GetAdmin() error {
	err := db.GetAdmin()

	if err != nil {
		return kerr.E(err, kerr.SecureKind)
	}

	return nil
}
