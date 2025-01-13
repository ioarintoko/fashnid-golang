package handler

import (
	"database/sql"
	comments "fashnid/handler/Comments"
	favorites "fashnid/handler/Favorites"
	loves "fashnid/handler/Loves"
	posts "fashnid/handler/Posts"
	profiles "fashnid/handler/Profiles"
	users "fashnid/handler/Users"
	"fashnid/handler/authentications"

	"fmt"
	"net/http"
	"strings"
)

var DB *sql.DB

func RegisDB(db *sql.DB) {
	DB = db
}

const (
	login		=	"login"
	logout 		= 	"logout"
	postdata	= 	"postdata"
	upload  	= 	"upload"
	profile 	= 	"profile"
	love		=	"love"
	user		=	"user"
	comment		=	"comment"
	favorite	=	"favorite"
)

func API(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Header", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	url := r.URL.Path
	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")

	fmt.Println("token center", w.Header()["Set-Cookie"])
	
	switch dataURL[2] {
		case login:
			authentications.Auth(DB, w, r)
		
		case logout:
			authentications.Auth(DB, w, r)
		
		case profile:
			profiles.Profiles(DB, w, r)

		case postdata:
			posts.PostDatas(DB, w, r)

		case love:
			loves.Loves(DB, w, r)

		case user:
			users.Users(DB, w, r)

		case favorite:
			favorites.Favorites(DB, w, r)

		case comment:
			comments.Comments(DB, w, r)
			
		default:
			fmt.Println("Wrong Path")
			
	}
}