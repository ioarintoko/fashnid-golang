package loves

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

func Gets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("params")
	datalove, err := model.GetsLove(db, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	jsonData, err := json.Marshal(datalove)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}

func Get(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// cookie := tokenize.Decode(w, r)
	// fmt.Println(reflect.TypeOf(int(cookie["iduser"].(float64))))

	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]
	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// id := int(cookie["iduser"].(float64))
	// fmt.Printf("idnya %v", id)
	datalove := &model.Love{IDLove: id}
	err = datalove.Get(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(datalove)

	jsonData, err := json.Marshal(datalove)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}

func Delete(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if lastIndex != "love" {
		datalove := model.Love{IDLove: id}
		err := datalove.Delete(db)
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

	if lastIndex == "love" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		var love model.Love
		err = json.Unmarshal(body, &love)
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
		love.IDUser = iduser

		err = love.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex != "love" {
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

		datalove := model.Love{IDLove: id}
		err = datalove.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("OK"))
		}
	}
}