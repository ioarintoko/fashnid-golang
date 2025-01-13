package test

import (
	"fashnid/model"
	"fmt"
	"testing"
)

var datacom = []*model.Comment{
	{
		IDComment: 1,
		IDPost: 1,
		IDUser: 1,
		Comment: "dssjkfnsdnjsfndsn",
	},
	{
		IDComment: 2,
		IDPost: 1,
		IDUser: 1,
		Comment: "dssjkfnsdnjsfndsn",
	},
	{
		IDComment: 3,
		IDPost: 1,
		IDUser: 1,
		Comment: "dssjkfnsdnjsfndsn",
	},
}

func TestComment(t *testing.T){
	t.Run("Test Insert Table Comment", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datacom{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Comment", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"idpost": 3,
		}

		err := datacom[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Comment", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datacom := model.Comment{IDComment: 2}
		err := datacom.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Comment", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datacom := model.Comment{IDComment: 1}
		err := datacom.Get(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets Table Comment", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datacom, err := model.GetsComment(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datacom {
			fmt.Println(*val)
		}
	})
}