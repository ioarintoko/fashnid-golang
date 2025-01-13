package model

import (
	"database/sql"
	"fashnid/lib"
	"fmt"
	"strings"
)

type Post struct {
	IDPost 		int    		`json:"idpost"`
	IDUser		int			`json:"iduser"`
	Post   		string 		`json:"post"`
	Photo   	string 		`json:"photo"`
	Username	string		`json:"username"`
	SumLike		int			`json:"sumlike"`
}

var TablePost = lib.Table{
	Name: "Post",
	Field: []string{
		"IDPost INT(9) AUTO_INCREMENT PRIMARY KEY",
		"IDUser INT(9)",
		"Post TEXT",
	},
}

var ForeignKeyPostUser = lib.ForeignKey{
	Name: "Post",
	ForeignName: "User",
	Field: "IDUser",
	ForeignField: "IDUser",
}

func(p *Post) Insert(db *sql.DB) error {
	query := "INSERT INTO Post (IDUser, Post) VALUES(?,?)"
	_, err := db.Exec(query, p.IDUser, p.Post)
	return err
}

func(p *Post) Delete(db *sql.DB) error {
	query := "DELETE FROM User WHERE IDUser = ?"
	_, err := db.Exec(query, p.IDPost)
	return err
}

func (p *Post) Update(db *sql.DB, datapost map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	for key, value := range datapost {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Post SET %s WHERE IDPost = %d", dataUpdate, p.IDPost)
	_, err := db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (p *Post) Get(db *sql.DB) (*Post, error) {
	query := `SELECT Post.*, User.Username, Count(Love.IDPost) FROM Post
					INNER JOIN Love ON Post.IDPost = Love.IDPost
					INNER JOIN User ON User.IDUser = Post.IDUser
					WHERE Post.IDPost = ?`
	datapost, err := db.Query(query, p.IDPost)
	if err != nil {
		return nil, err
	}
	defer datapost.Close()

	var result *Post
	for datapost.Next(){
		each := &Post{}
		err = datapost.Scan(&each.IDPost, &each.IDUser, &each.Post, &each.Username, &each.SumLike)
		if err != nil {
			return nil, err
		}
		result = each
	}

	return result, nil
}

func (p *Post) GetByUserID(db *sql.DB) ([]*Post, error) {
	query := `SELECT Photo.photo FROM Post 
				INNER JOIN Photo ON Post.IDPost = Photo.IDPost
				WHERE Post.IDUser = ?`
	datauser, err := db.Query(query, &p.IDUser)
	if err != nil {
		return nil, err
	}
	defer datauser.Close()

	var result []*Post
	for datauser.Next(){
		each := &Post{}
		err = datauser.Scan(&each.Photo)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	return result, nil
}

func GetsPost(db *sql.DB, params ...string) ([]*Post, error) {
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
	query = `SELECT Post.*, User.Username, (SELECT COUNT(Love.IDPost) FROM love WHERE love.IDPost = Post.IDPost) as sumLike
			 FROM Post 
				INNER JOIN User ON Post.IDUser = User.IDUser`
	if dataKondisi != "" {
		query += " WHERE " + dataKondisi
	}

	datapost, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datapost.Close()

	var result []*Post
	for datapost.Next() {
		each := &Post{}
		err := datapost.Scan(&each.IDPost, &each.IDUser, &each.Post, &each.Username, &each.SumLike)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}