package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Follow struct {
	IDFollow		int    	`json:"idfollow"`
	IDUserFlwg		int		`json:"iduserflwg"`
	IDUserFlwd		int		`json:"iduserflwd"`
	Status			string	`json:"status"`
}

var TableFollow = lib.Table{
	Name: "Follow",
	Field: []string{
		"IDFollow INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDUserFlwg INT(9)",
		"IDUserFlwd INT(9)",
		"Status VARCHAR(20)",
	},
}

var ForeignKeyFolUser1 = lib.ForeignKey{
	Name: "Follow",
	ForeignName: "User",
	Field: "IDUserFlwd",
	ForeignField: "IDUser",
}

var ForeignKeyFolUser2 = lib.ForeignKey{
	Name: "Follow",
	ForeignName: "User",
	Field: "IDUserFlwg",
	ForeignField: "IDUser",
}

func(fo *Follow) Insert(db *sql.DB) error {
	query := "INSERT INTO Follow (IDUserFlwg, IDUserFlwd, Status) VALUES(?,?,?)"
	_, err := db.Exec(query, fo.IDFollow, fo.IDUserFlwg, fo.IDUserFlwd, fo.Status)
	return err
}

func(fo *Follow) Delete(db *sql.DB) error {
	query := "DELETE FROM Follow WHERE IDFollow = ?"
	_, err := db.Exec(query, fo.IDFollow)
	return err
}

func (fo *Follow) Update(db *sql.DB, datafo map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range datafo {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Follow SET %s WHERE IDFollow = %d", dataUpdate, fo.IDFollow)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (fo *Follow) Get(db *sql.DB) error {
	query := "SELECT * FROM Follow WHERE IDFollow = ?"
	err := db.QueryRow(query, fo.IDFollow).Scan(&fo.IDFollow, &fo.IDUserFlwg, &fo.IDUserFlwd, &fo.Status)
	return err
}

func GetsFollow(db *sql.DB, params ...string) ([]*Follow, error) {
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
	query = "SELECT * FROM Follow"
	if dataKondisi != "" {
		query = "SELECT * FROM Follow WHERE " + dataKondisi
	}

	datafo, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datafo.Close()

	var result []*Follow
	for datafo.Next() {
		each := &Follow{}
		err := datafo.Scan(&each.IDFollow, &each.IDUserFlwg, &each.IDUserFlwd, &each.Status)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}