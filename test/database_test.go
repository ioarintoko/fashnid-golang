package test

import (
	"fashnid/lib"
	"fashnid/model"
	"testing"
)

var database, databaseDefaultMysql string

func init() {
	database = "fashnid"
	databaseDefaultMysql = "Mysql"
}

func TestDatabaseMysql(t *testing.T) {
	t.Run("MySQL Connection Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
	})

	t.Run("Drop Table Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.DropDB(db, database)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create DB Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateDB(db, database)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Post", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TablePost)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Photo", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TablePhoto)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Favorite", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableFav)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Like", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableLove)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Comment", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableComment)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Follow", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableFollow)
		if err != nil {
			t.Fatal(err)
		}
	})

	//Foreign Key
	t.Run("Test Add Foreign Key Post - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyPostUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Photo - Post", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyPhotoPost)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Love - Post", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyLovePost)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Love - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyLoveUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Fav - Post", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyFavPost)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Fav - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyFavUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Comment - Post", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyCommPost)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Comment - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyCommUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Following - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyFolUser1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Followed - User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyFolUser2)
		if err != nil {
			t.Fatal(err)
		}
	})
}