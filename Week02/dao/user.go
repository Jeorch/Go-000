package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	ID 		int
	Name    string
	Age     int
}

func (u *User) findUserById(id int) error {
	err := db.QueryRow("SELECT * FROM USERS WHERE ID = ?", id).Scan(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		} else {
			return errors.Wrap(err, "undefined error")
		}
	}
	return nil
}
