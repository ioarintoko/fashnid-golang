package middleware

import "database/sql"

type Login struct {
	IDUser		int	`json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func(lgn *Login) Login(db *sql.DB) error {
	query := "SELECT IDUser FROM User WHERE Username = ? AND Password = MD5(?)"
	err := db.QueryRow(query, &lgn.Username, &lgn.Password).Scan(&lgn.IDUser)
	return err
}