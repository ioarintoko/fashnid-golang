package favorites

import (
	"database/sql"
	"encoding/json"

	"fashnid/model"
	"fashnid/tokenize"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// func Gets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query().Get("params")
// 	datacom, err := model.GetsComment(db, params)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	jsonData, err := json.Marshal(datacom)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	w.Write(jsonData)
// }

// func Get(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// cookie := tokenize.Decode(w, r)
// 	// fmt.Println(reflect.TypeOf(int(cookie["iduser"].(float64))))

// 	url := r.URL.Path
// 	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
// 	lastIndex := dataUrl[len(dataUrl)-1]
// 	id, err := strconv.Atoi(lastIndex)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	// id := int(cookie["iduser"].(float64))
// 	// fmt.Printf("idnya %v", id)
// 	datacom := &model.Comment{IDComment: id}
// 	err = datacom.Get(db)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	fmt.Println(datauser)

// 	jsonData, err := json.Marshal(datauser)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	w.Write(jsonData)
// }

func Delete(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var cookie jwt.MapClaims
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}

	if token != nil {
		cookie = tokenize.Decode(w, r)
	}
	
	iduser := int(cookie["iduser"].(float64))

	if lastIndex != "favorite" {
		datafav := model.Favorite{IDPost: id, IDUser: iduser}
		err := datafav.Delete(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte("ok"))
	}
}

func Insert(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex == "favorite" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		var favorite model.Favorite
		err = json.Unmarshal(body, &favorite)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	
		var cookie jwt.MapClaims
		token, err := r.Cookie("token")
		if err != nil {
			fmt.Println(err)
		}
	
		if token != nil {
			cookie = tokenize.Decode(w, r)
		}
		
		iduser := int(cookie["iduser"].(float64))
		favorite.IDUser = iduser
		err = favorite.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex != "favorite" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Println(body)
		fmt.Println(jsonMap)

		id, err := strconv.Atoi(lastIndex)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		datafav := model.Favorite{IDFav: id}
		err = datafav.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("OK"))
		}
	}
}