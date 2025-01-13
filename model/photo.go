package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Photo struct {
	IDPhoto		int    	`json:"idphoto"`
	IDPost		int		`json:"idpost"`
	Photo		string	`json:"photo"`
	Caption		string	`json:"caption"`
	Link		string	`json:"link"`
}

var TablePhoto = lib.Table{
	Name: "Photo",
	Field: []string{
		"IDPhoto INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDPost INT(9)",
		"Photo TEXT",
		"Caption VARCHAR(35)",
		"Link TEXT",
	},
}

var ForeignKeyPhotoPost = lib.ForeignKey{
	Name: "Photo",
	ForeignName: "Post",
	Field: "IDPost",
	ForeignField: "IDPost",
}

func(ph *Photo) Insert(db *sql.DB) error {
	query := "INSERT INTO Photo (IDPost, Photo, Caption, Link) VALUES(?,?,?,?)"
	_, err := db.Exec(query, ph.IDPost, ph.Photo, ph.Caption, ph.Link)
	return err
}

func(ph *Photo) Delete(db *sql.DB) error {
	query := "DELETE FROM Photo WHERE IDPhoto = ?"
	_, err := db.Exec(query, ph.IDPhoto)
	return err
}

func (ph *Photo) Update(db *sql.DB, dataph map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range dataph {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Photo SET %s WHERE IDPhoto = %d", dataUpdate, ph.IDPhoto)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (ph *Photo) Get(db *sql.DB) []*Photo {
	query := "SELECT * FROM Photo WHERE IDPost = ?"
	dataph, err := db.Query(query, ph.IDPost)
	if err != nil {
		fmt.Println(err)
	}

	var result []*Photo
	for dataph.Next() {
		each := &Photo{}
		err := dataph.Scan(&each.IDPhoto, &each.IDPost, &each.Photo, &each.Caption, &each.Link)
		if err != nil {
			fmt.Println(err)
		}
		
		result = append(result, each)
	}

	return result
}

func GetsPhoto(db *sql.DB, params ...string) ([]*Photo, error) {
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
	query = "SELECT * FROM Photo"
	if dataKondisi != "" {
		query = "SELECT * FROM Photo WHERE " + dataKondisi
	}

	dataph, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer dataph.Close()

	var result []*Photo
	for dataph.Next() {
		each := &Photo{}
		err := dataph.Scan(&each.IDPhoto, &each.IDPost, &each.Photo, &each.Caption, &each.Link)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}