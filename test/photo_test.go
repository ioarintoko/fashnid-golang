package test

import (
	"fashnid/model"
	"fmt"
	"testing"
)

var dataph = []*model.Photo{
	{
		IDPhoto: 1,
		IDPost: 1,
		Photo: "picture",
		Caption: "kjfnwjfnjkf",
		Link: "fjndfjndjfdf",
	},
	{
		IDPhoto: 2,
		IDPost: 3,
		Photo: "picture",
		Caption: "kjfnwjfnjkf",
		Link: "fjndfjndjfdf",
	},
	{
		IDPhoto: 3,
		IDPost: 1,
		Photo: "picture",
		Caption: "kjfnwjfnjkf",
		Link: "fjndfjndjfdf",
	},
}

func TestPhoto(t *testing.T){
	t.Run("Test Insert Table Photo", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range dataph{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Photo", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"idpost": 3,
		}

		err := dataph[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Photo", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataph := model.Photo{IDPhoto: 2}
		err := dataph.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Photo", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataph := model.Photo{IDPhoto: 1}
		err := dataph.Get(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets Table Photo", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataph, err := model.GetsPhoto(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range dataph {
			fmt.Println(*val)
		}
	})
}