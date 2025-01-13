package comments

import (
	"database/sql"
	"encoding/json"
	"fashnid/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

	if lastIndex != "comment" {
		datacom := model.Comment{IDComment: id}
		err := datacom.Delete(db)
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

	if lastIndex == "comment" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		var comment model.Comment
		err = json.Unmarshal(body, &comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = comment.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex != "comment" {
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

		datacom := model.Comment{IDComment: id}
		err = datacom.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("OK"))
		}
	}
}