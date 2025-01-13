package test

import (
	"fashnid/model"
	"fmt"
	"testing"
)

var datafav = []*model.Favorite{
	{
		IDFav: 1,
		IDPost: 1,
		IDUser: 1,
	},
	{
		IDFav: 2,
		IDPost: 3,
		IDUser: 1,
	},
}

func TestFav(t *testing.T){
	t.Run("Test Insert Table Favorite", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datafav{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Favorite", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"idpost": 3,
		}

		err := datafav[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Favorite", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datafav := model.Favorite{IDFav: 2}
		err := datafav.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Favorite", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datafav := model.Favorite{IDFav: 1}
		err := datafav.Get(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets Table Favorite", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datafav, err := model.GetsFav(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datafav {
			fmt.Println(*val)
		}
	})
}