package test

import (
	"database/sql"
	"fashnid/lib"
	"fashnid/model"
	"fmt"
	"testing"
)

var datauser = []*model.User{
	{
		IDUser: 1,
		Name: "Bramantio",
		Email: "Bmantio20@gmail.com",
		Username: "Bram7",
		Password: "121212",
		Bio: "lorem ipsum dolor sit amet",
		Photo: "picture",
	},
	{
		IDUser: 2,
		Name: "Galih",
		Email: "bamzz.bamz@yahoo.com",
		Username: "Bram1",
		Password: "121212",
		Bio: "lorem ipsum dolor sit amet",
		Photo: "picture",
	},
}

func ConnectDB(t *testing.T) (*sql.DB, error){
	db, err := lib.ConnectMySql(database)
	if err != nil {
		t.Fatal(err)
		return nil, err
	}
	
	return db, nil
}

func TestUser(t *testing.T){
	t.Run("Test Insert Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _,val := range datauser{
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{} {
			"name" : "Bramantonio",
		}

		err := datauser[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datauser := model.User{IDUser: 2}
		err := datauser.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datauser := model.User{IDUser: 1}
		err := datauser.Get(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(datauser)
	})

	t.Run("Test Gets Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datauser, err := model.GetsUser(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datauser{
			fmt.Println(*val)
		}
	})
}