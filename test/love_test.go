package test

import (
	"fashnid/model"
	"fmt"
	"testing"
)

var datalove = []*model.Love{
	{
		IDLove: 1,
		IDPost: 1,
		IDUser: 1,
	},
	{
		IDLove: 2,
		IDPost: 3,
		IDUser: 1,
	},
}

func TestLove(t *testing.T){
	t.Run("Test Insert Table Love", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datalove{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Love", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"idpost": 3,
		}

		err := datalove[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Love", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datalove := model.Love{IDLove: 2}
		err := datalove.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Love", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datalove := model.Love{IDLove: 1}
		err := datalove.Get(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets Table Love", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datalove, err := model.GetsLove(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datalove {
			fmt.Println(*val)
		}
	})
}