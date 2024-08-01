package service

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lucasmb/getth-cms/models"
)


func NewPostService(db *sqlx.DB) *PostService {
	return &PostService{
		Post: models.Post{},
		Db: db,
	}
}

type PostService struct {
	Post	models.Post
	Db *sqlx.DB
}

func (ps *PostService) List(status string) ([]models.Post, error) {
	query := "SELECT id, title, content, author_id, status, created_at FROM posts WHERE status = ?;"
	
	rows, err := ps.Db.Queryx(query, status)
	if err != nil {
		return []models.Post{}, err 
	}

	posts := []models.Post{}
	for rows.Next() {
		var p models.Post
		err = rows.StructScan(&p)
		if err != nil {
			log.Fatalln(err)
		} 
		posts = append(posts, p)

	}

	return posts, nil

}