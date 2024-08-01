package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


type User struct {
	Id int `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string 
	Email     string `db:"email"`
	Password  string `db:"password"`
	Role      string 
	Active    bool   `db:"active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	TenantId  int    `db:"tenant_id"`
}

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./db/echo.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}