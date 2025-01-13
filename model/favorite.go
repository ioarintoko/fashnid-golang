package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Favorite struct {
	IDFav	int		`json:"idfav"`
	IDPost 	int    	`json:"idpost"`
	IDUser  int 	`json:"iduser"`
	Photo	string	`json:"photo"`
	Isfav	int		`json:"isfav"`
}

var TableFav = lib.Table{
	Name: "Favorite",
	Field: []string{
		"IDFav INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDPost INT(9)",
		"IDUser INT(9)",
	},
}

var ForeignKeyFavPost = lib.ForeignKey{
	Name: "Favorite",
	ForeignName: "Post",
	Field: "IDPost",
	ForeignField: "IDPost",
}

var ForeignKeyFavUser = lib.ForeignKey{
	Name: "Favorite",
	ForeignName: "User",
	Field: "IDUser",
	ForeignField: "IDUser",
}

func(f *Favorite) Insert(db *sql.DB) error {
	query := "INSERT INTO Favorite (IDPost, IDUser) VALUES(?,?)"
	_, err := db.Exec(query, f.IDPost, f.IDUser)
	return err
}

func(f *Favorite) Delete(db *sql.DB) error {
	query := "DELETE FROM Favorite WHERE IDPost = ? AND IDUser = ?"
	_, err := db.Exec(query, f.IDPost, f.IDUser)
	return err
}

func (f *Favorite) Update(db *sql.DB, datafav map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range datafav {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Favorite SET %s WHERE IDFav = %d", dataUpdate, f.IDFav)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func(f *Favorite) GetIsFav(db *sql.DB) bool {
	query := "SELECT COUNT(IDUser) FROM Favorite WHERE IDUser = ? AND IDPost = ?"
	_ = db.QueryRow(query, f.IDUser, f.IDPost).Scan(&f.Isfav)
	var isfav bool
	if f.Isfav > 0 {
		isfav = true
	} else {
		isfav = false
	}
	return isfav
}

func(f *Favorite) Get(db *sql.DB) error {
	query := "SELECT * FROM Favorite WHERE IDFav = ?"
	err := db.QueryRow(query, f.IDFav).Scan(&f.IDFav, &f.IDPost, &f.IDUser)
	return err
}

func(f *Favorite) GetByUserID(db *sql.DB) ([]*Favorite, error) {
	query := `SELECT Photo.photo FROM Favorite 
				INNER JOIN Post ON Favorite.IDPost = Post.IDPost
				INNER JOIN Photo ON Post.IDPost = Photo.IDPost
				WHERE Favorite.IDUser = ?`
	datauser, err := db.Query(query, &f.IDUser)
	if err != nil {
		return nil, err
	}
	defer datauser.Close()

	var result []*Favorite
	for datauser.Next(){
		each := &Favorite{}
		err = datauser.Scan(&each.Photo)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	return result, nil
}

func GetsFav(db *sql.DB, params ...string) ([]*Favorite, error) {
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
	query = "SELECT * FROM Favorite"
	if dataKondisi != "" {
		query = "SELECT * FROM Favorite WHERE " + dataKondisi
	}

	datafav, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datafav.Close()

	var result []*Favorite
	for datafav.Next() {
		each := &Favorite{}
		err := datafav.Scan(&each.IDFav, &each.IDPost, &each.IDUser)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}