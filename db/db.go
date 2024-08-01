package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)


type Store struct {
	Db *sqlx.DB
}

func NewStore(dbName string) (Store, error) {
	Db, err := getConnection(dbName)
	if err != nil {
		return Store{}, err
	}

	return Store{
		Db,
	}, nil
}

func getConnection(dbName string) (*sqlx.DB, error) {
	var (
		err error
		db  *sqlx.DB
	)

	if db != nil {
		return db, nil
	}

	// Init SQLite3 database
	db, err = sqlx.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}

// TODO: replace with store
func InitDb(dbFile string) *sqlx.DB {

	db, err := sqlx.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//seedDb(db)

	return db
}