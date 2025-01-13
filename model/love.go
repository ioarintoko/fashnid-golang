package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Love struct {
	IDLove	int		`json:"idlove"`
	IDPost 	int    	`json:"idpost"`
	IDUser  int 	`json:"iduser"`
	Islove	int		`json:"islove"`
}

var TableLove = lib.Table{
	Name: "Love",
	Field: []string{
		"IDLove INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDPost INT(9)",
		"IDUser INT(9)",
	},
}

var ForeignKeyLovePost = lib.ForeignKey{
	Name: "Love",
	ForeignName: "Post",
	Field: "IDPost",
	ForeignField: "IDPost",
}

var ForeignKeyLoveUser = lib.ForeignKey{
	Name: "Love",
	ForeignName: "User",
	Field: "IDUser",
	ForeignField: "IDUser",
}

func(lk *Love) Insert(db *sql.DB) error {
	query := "INSERT INTO Love (IDPost, IDUser) VALUES(?,?)"
	_, err := db.Exec(query, lk.IDPost, lk.IDUser)
	return err
}

func(lk *Love) Delete(db *sql.DB) error {
	query := "DELETE FROM Love WHERE IDLove = ?"
	_, err := db.Exec(query, lk.IDLove)
	return err
}

func (lk *Love) Update(db *sql.DB, dataLove map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range dataLove {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Love SET %s WHERE IDLove = %d", dataUpdate, lk.IDLove)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func(lk *Love) GetIsLove(db *sql.DB) bool {
	query := "SELECT COUNT(IDUser) FROM Love WHERE IDUser = ? AND IDPost = ?"
	_ = db.QueryRow(query, lk.IDUser, lk.IDPost).Scan(&lk.Islove)
	var islove bool
	if lk.Islove > 0 {
		islove = true
	} else {
		islove = false
	}
	return islove
}

func(lk *Love) Get(db *sql.DB) error {
	query := "SELECT * FROM Love WHERE IDLove = ?"
	err := db.QueryRow(query, lk.IDLove).Scan(&lk.IDLove, &lk.IDPost, &lk.IDUser)
	return err
}

func GetsLove(db *sql.DB, params ...string) ([]*Love, error) {
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
	query = "SELECT * FROM Love"
	if dataKondisi != "" {
		query = "SELECT * FROM Love WHERE " + dataKondisi
	}

	dataLove, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer dataLove.Close()

	var result []*Love
	for dataLove.Next() {
		each := &Love{}
		err := dataLove.Scan(&each.IDLove, &each.IDPost, &each.IDUser)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}