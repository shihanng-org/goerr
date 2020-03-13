package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

func GetUser() error {
	return errors.Wrap(sql.ErrNoRows, "db: get user")
}

func GetAdmin() error {
	return errors.Wrap(sql.ErrNoRows, "db: get admin")
}
