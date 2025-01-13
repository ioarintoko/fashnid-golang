package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Comment struct {
	IDComment	int    	`json:"idcomment"`
	IDPost		int		`json:"idpost"`
	IDUser		int		`json:"iduser"`
	Comment		string	`json:"comment"`
}

var TableComment = lib.Table{
	Name: "Comment",
	Field: []string{
		"IDComment INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDPost INT(9)",
		"IDUser INT(9)",
		"Comment TEXT",
	},
}

var ForeignKeyCommPost = lib.ForeignKey{
	Name: "Comment",
	ForeignName: "Post",
	Field: "IDPost",
	ForeignField: "IDPost",
}

var ForeignKeyCommUser = lib.ForeignKey{
	Name: "Comment",
	ForeignName: "User",
	Field: "IDUser",
	ForeignField: "IDUser",
}

func(c *Comment) Insert(db *sql.DB) error {
	query := "INSERT INTO Comment (IDPost, IDUser, Comment) VALUES(?,?,?)"
	_, err := db.Exec(query, c.IDPost, c.IDUser, c.Comment)
	return err
}

func(c *Comment) Delete(db *sql.DB) error {
	query := "DELETE FROM Comment WHERE IDComment = ?"
	_, err := db.Exec(query, c.IDComment)
	return err
}

func (c *Comment) Update(db *sql.DB, datacom map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range datacom {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Comment SET %s WHERE IDComment = %d", dataUpdate, c.IDComment)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (c *Comment) Get(db *sql.DB) []*Comment {
	query := "SELECT * FROM Comment WHERE IDPost = ?"
	datacom, err := db.Query(query, c.IDPost)
	if err != nil {
		fmt.Println(err)
	}
	
	var result []*Comment
	for datacom.Next() {
		each := &Comment{}
		err := datacom.Scan(&each.IDComment, &each.IDPost, &each.IDUser, &each.Comment)
		if err != nil {
			fmt.Println(err)
		}

		result = append(result, each)
	}

	return result
}

func GetsComment(db *sql.DB, params ...string) ([]*Comment, error) {
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
	query = "SELECT * FROM Comment"
	if dataKondisi != "" {
		query = "SELECT * FROM Comment WHERE " + dataKondisi
	}

	datacom, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datacom.Close()

	var result []*Comment
	for datacom.Next() {
		each := &Comment{}
		err := datacom.Scan(&each.IDComment, &each.IDPost, &each.IDUser, &each.Comment)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}