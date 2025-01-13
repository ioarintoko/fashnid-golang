package services

import (
	"database/sql"
	"fashnid/model"
	"fmt"
)

type Profile struct {
	User 		*model.User
	Post		[]*model.Post
	Fav 		[]*model.Favorite
	IsOwner		bool				`json:"isowner"`
}

func(pr *Profile) Get(db *sql.DB, idVisit int, idOwner int) (*Profile, error) {
	userorgn := model.User{IDUser: idVisit}
	datauser, err := userorgn.Profile(db)
	if err != nil {
		return nil, err
	}

	post := model.Post{IDUser: idVisit}
	datapost, err := post.GetByUserID(db)
	if err != nil {
		fmt.Println(err)
	}

	fav := model.Favorite{IDUser: idVisit}
	datafav, err := fav.GetByUserID(db)
	if err != nil {
		fmt.Println(err)
	}

	isOwner := false
	if(idOwner == idVisit){
		isOwner = true
	}

	data := &Profile{
		User: datauser,
		Post: datapost,
		Fav: datafav,
		IsOwner: isOwner,
	}

	return data, nil
}