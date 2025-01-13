package favorites

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Route(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// url := r.URL.Path
	// dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	// lastIndex := dataUrl[len(dataUrl)-1]
	switch r.Method{
		case http.MethodPost:
			Insert(db, w, r)
		
		case http.MethodPut:
			Update(db, w, r)

		case http.MethodDelete:
			Delete(db, w, r)
		
		default:
			fmt.Println("Wrong Method Admin")
	}
}