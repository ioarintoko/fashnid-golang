package loves

import (
	"database/sql"
	"net/http"
)

func Loves(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}