package comments

import (
	"database/sql"
	"net/http"
)

func Comments(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}