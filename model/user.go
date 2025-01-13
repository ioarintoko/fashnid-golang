package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type User struct {
	IDUser   int    `json:"iduser"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	Photo    string `json:"photo"`
}

var TableUser = lib.Table{
	Name: "User",
	Field: []string{
		"IDUser INT(9) AUTO_INCREMENT PRIMARY KEY",
		"Name VARCHAR(50)",
		"Email VARCHAR(100)",
		"Username VARCHAR(15)",
		"Password TEXT",
		"Bio TEXT",
		"Photo TEXT",
	},
}

func(u *User) Insert(db *sql.DB) error {
	query := `INSERT INTO User (Name, Email, Username, Password, Bio, Photo)
				VALUES(?,?,?,MD5(?),?,?)`

	_, err := db.Exec(query, u.Name, u.Email, u.Username, u.Password, u.Bio, u.Photo)
	return err
}

func(u *User) Delete(db *sql.DB) error {
	query := "DELETE FROM User WHERE IDUser = ?"
	_, err := db.Exec(query, u.IDUser)
	return err
}

func (u *User) Update(db *sql.DB, datauser map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range datauser {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE User SET %s WHERE IDUser = %d", dataUpdate, u.IDUser)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (u *User) Get(db *sql.DB) error {
	query := "SELECT * FROM User WHERE IDUser = ?"
	err := db.QueryRow(query, u.IDUser).Scan(&u.IDUser, &u.Name, &u.Email, &u.Username, &u.Password, 
		&u.Bio, &u.Photo)
	return err
}

func(u *User) Profile(db *sql.DB) (*User, error){
	query := "SELECT * FROM User WHERE IDUser = ?"
	datauser, err := db.Query(query, &u.IDUser)
	if err != nil {
		return nil, err
	}
	defer datauser.Close()

	var result *User
	for datauser.Next(){
		each := &User{}
		err = datauser.Scan(&each.IDUser, &each.Name, &each.Email, &each.Username, &each.Password, &each.Bio, &each.Photo)
		if err != nil {
			return nil, err
		}
		result = each
	}

	return result, nil
}

func GetsUser(db *sql.DB, params ...string) ([]*User, error) {
	var kolom = []string{}
	var args []interface{}

	if len(params) != 0 {
		if params[0] != "" {
			dataParams := strings.Split(params[len(params)-1], ";")
			for _, val := range dataParams {
				temp := strings.Split(fmt.Sprintf("%s", val), ",")
				where := fmt.Sprintf("%s %s ?", strings.ToLower(temp[0]), temp[1])
				kolom = append(kolom, where)
				args = append(args, temp[2])
			}
		}
	}

	dataKondisi := strings.Join(kolom, " AND ")
	var query string
	query = "SELECT * FROM User"
	if dataKondisi != "" {
		query = "SELECT * FROM User WHERE " + dataKondisi
	}

	datauser, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datauser.Close()

	var result []*User
	for datauser.Next() {
		each := &User{}
		err := datauser.Scan(&each.IDUser, &each.Name, &each.Email, &each.Username, 
			&each.Password, &each.Bio, &each.Photo)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}