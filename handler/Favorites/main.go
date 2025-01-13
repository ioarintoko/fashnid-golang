package favorites

import (
	"database/sql"
	"net/http"
)

func Favorites(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}