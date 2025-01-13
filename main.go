package main

import (
	"fashnid/handler"
	"fashnid/lib"
	"fmt"
	"net/http"
)

var database string
func init(){
	database = "fashnid"
}

func main() {
	db, err := lib.ConnectMySql(database)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	
	handler.RegisDB(db)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("assets/picture/"))))
	http.HandleFunc("/api/", handler.API)
	http.ListenAndServe(":8087", nil)
}