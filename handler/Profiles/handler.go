package profiles

import (
	"database/sql"
	"encoding/json"

	"fashnid/services"
	"fashnid/tokenize"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Get(db *sql.DB,w http.ResponseWriter, r *http.Request) {
	var cookie jwt.MapClaims
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}

	if token != nil {
		cookie = tokenize.Decode(w, r)
	}
	
	var iduser int
	if cookie != nil {
		iduser = int(cookie["iduser"].(float64))
	} else {
		iduser=0
	}
	

	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]
	idvisit, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	profile := &services.Profile{}
	dataProfile, err := profile.Get(db, idvisit, iduser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataProfile)

	jsonData, err := json.Marshal(dataProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}