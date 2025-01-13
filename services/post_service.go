package services

import (
	"context"
	"database/sql"
	"fashnid/model"
	"fmt"
	"log"
)

type PostData struct {
	Photo 		[]*model.Photo
	Comments	[]*model.Comment
	Post		*model.Post
	Isfav		bool	`json:"isfav"`
	Islove		bool	`json:"islove"`
}

func(pd *PostData) Delete(db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	queryPost := `DELETE FROM Photo WHERE IDPost = ?`

	_, err = tx.Exec(queryPost, pd.Post.Post)
	if err != nil {
		tx.Rollback()
	}

	queryPhoto := `DELETE FROM Post WHERE IDPost = ?`

	for i:=0;i<len(pd.Photo);i++ {
		_, err := tx.Exec(queryPhoto, pd.Post.IDPost)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	
	return err
}

func(pd *PostData) Insert(db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	queryPost := `INSERT INTO Post (IDUser, Post) VALUES (?,?)`

	response, err := tx.Exec(queryPost, pd.Post.IDUser, pd.Post.Post)
	if err != nil {
		tx.Rollback()
	}

	idPost, err := response.LastInsertId()
	if err != nil {
		tx.Rollback()
	}

	queryPhoto := `INSERT INTO Photo (IDPost, Photo, Caption, Link) VALUES (?,?,?,?)`

	for i:=0;i<len(pd.Photo);i++ {
		_, err := tx.Exec(queryPhoto, idPost, pd.Photo[i].Photo, pd.Photo[i].Caption, pd.Photo[i].Link)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	
	return err
}

func GetsPostData(db *sql.DB, idUser int) []*PostData {
	datapost, err := model.GetsPost(db)
	if err != nil {
		fmt.Println(err)
	}

	islove := false
	isfav := false

	// if idUser != 0 {
	// 	loveorgn := model.Love{IDPost: idPost, IDUser: idUser}
	// 	datalove := loveorgn.GetIsLove(db)
	// 	islove = datalove

	// 	favorgn := model.Favorite{IDPost: idPost, IDUser: idUser}
	// 	datafav := favorgn.GetIsFav(db)
	// 	isfav = datafav
	// }

	var postData []*PostData
	for i:=0; i < len(datapost); i++ {
		if idUser != 0 {
			loveorgn := model.Love{IDPost: datapost[i].IDPost, IDUser: idUser}
			datalove := loveorgn.GetIsLove(db)
			islove = datalove
			fmt.Println(islove)

			favorgn := model.Favorite{IDPost: datapost[i].IDPost, IDUser: idUser}
			datafav := favorgn.GetIsFav(db)
			isfav = datafav
			fmt.Println(isfav)
		}

		photoorgn := model.Photo{IDPost: datapost[i].IDPost}
		dataph := photoorgn.Get(db)

		data := &PostData{
			Post: datapost[i],
			Photo: dataph,
			Isfav: isfav,
			Islove: islove,
		}
		postData = append(postData, data)
	}
	
	return postData
}

func(pd *PostData) Get(db *sql.DB, idPost int, idUser int) *PostData {
	params := fmt.Sprintf("idpost=%d", idPost)
	fmt.Println(params)

	postorgn := model.Post{IDPost: idPost}
	datapost, err := postorgn.Get(db)
	if err != nil {
		fmt.Println(err)
	}

	commentorgn := model.Comment{IDPost: idPost}
	datacom := commentorgn.Get(db)
	if err != nil {
		fmt.Println(err)
	}

	photoorgn := model.Photo{IDPost: idPost}
	dataph := photoorgn.Get(db)
	if err != nil {
		fmt.Println(err)
	}

	islove := false
	isfav := false

	if idUser != 0 {
		loveorgn := model.Love{IDPost: idPost, IDUser: idUser}
		datalove := loveorgn.GetIsLove(db)
		islove = datalove

		favorgn := model.Favorite{IDPost: idPost, IDUser: idUser}
		datafav := favorgn.GetIsFav(db)
		isfav = datafav
	}

	postData := &PostData{
		Photo: dataph,
		Comments: datacom,
		Post: datapost,
		Isfav: isfav,
		Islove: islove,
	}

	return postData
}