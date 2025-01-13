package test

import (
	"fashnid/model"
	"fmt"
	"testing"
)

var datapost = []*model.Post{
	{
		IDPost: 1,
		IDUser: 1,
		Post: "sdjkfjfnjskndsnfkjsdn",
	},
	{
		IDPost: 2,
		IDUser: 1,
		Post: "sdjkfjfnjskndsnfkjsdn",
	},
	{
		IDPost: 3,
		IDUser: 1,
		Post: "sdjkfjfnjskndsnfkjsdn",
	},
}

func TestPost(t *testing.T){
	t.Run("Test Insert Table Post", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datapost{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Post", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"post" : "denfjwkfnwjkfnwjnew",
		}

		err := datapost[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Post", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datapost := model.Post{IDPost: 2}
		err := datapost.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	// t.Run("Test Get Table Post", func(t *testing.T) {
	// 	db, _ := ConnectDB(t)
	// 	defer db.Close()

	// 	datapost := model.Post{IDPost: 1}
	// 	err := datapost.Get(db)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	t.Run("Test Gets Table Post", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datapost, err := model.GetsPost(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datapost {
			fmt.Println(*val)
		}
	})
}