package posts

import (
	"database/sql"
	"encoding/json"

	"fashnid/model"
	"fashnid/services"
	"fashnid/tokenize"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Delete(db *sql.DB, w http.ResponseWriter, r *http.Request){
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if lastIndex != "postdata" {
		dataprj := &services.PostData{Post: &model.Post{IDPost: id}}
		err := dataprj.Delete(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte("OK"))
	}
}

func Insert(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]
	
	if lastIndex == "post" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		var postData services.PostData
		err = json.Unmarshal(body, &postData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = postData.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func Gets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var cookie jwt.MapClaims
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Println("error handler gets",err)
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

	postDatas := services.GetsPostData(db, iduser)

	jsonData, err := json.Marshal(postDatas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}

func Get(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// cookie := tokenize.Decode(w, r)
	// iduser := int(cookie["iduser"].(float64))

	var cookie jwt.MapClaims
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Println("handler postdata",err)
	}

	if token != nil {
		cookie = tokenize.Decode(w, r)
	}
	
	var iduser int
	if cookie != nil {
		iduser = int(cookie["iduser"].(float64))
		fmt.Println(iduser)
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

	postdata := &services.PostData{}
	postData := postdata.Get(db, idvisit, iduser)
	
	fmt.Println(postData)

	jsonData, err := json.Marshal(postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}