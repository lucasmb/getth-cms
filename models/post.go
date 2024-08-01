package models

import "time"

//var DB *sql.DB

type Post struct {
		Id string `db:"id"`
		Title string `db:"title"`
		Content  string `db:"content"`
		AuthorId int `db:"author_id"`
		Email     string `db:"email"`
		Status      string 
		Active    bool   `db:"active"`
		CreatedAt time.Time `db:"created_at"` 
		UpdatedAt time.Time `db:"updated_at"`
		TenantId string `db:"tenant_id"`
}