package profiles

import (
	"database/sql"
	"net/http"
)

func Profiles(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}