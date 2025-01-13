package profiles

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Route(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Get(db, w, r)

	default:
		fmt.Println("Wrong Method Admin")
	}
}