package posts

import (
	"database/sql"
	"net/http"
)

func PostDatas(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}