package service

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/lucasmb/getth-cms/models"
)


type User struct {
	Id int `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string 
	Email     string `db:"email"`
	Password  string 
	Role      string 
	Active    bool   `db:"active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	TenantId string `db:"tenant_id"`
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		User: models.User{},
		Db: db,
	}
}

type UserService struct {
	User	models.User
	Db *sqlx.DB
}

func (us *UserService) Create(u models.User) (models.User, error) {

	query := `INSERT INTO users (first_name, last_name, username, email, password, role, tenant_id)
		VALUES(?, ?, ?, ?, ?, ?, ?) RETURNING *`

	stmt, err := us.Db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}

	err = stmt.QueryRow(
		u.FirstName,
		u.LastName,
		u.Username,
		u.Email,
		u.Password,
		u.Role,
		u.TenantId,
	).Scan(
		&us.User.Id,
		&us.User.FirstName,
		&us.User.LastName,
		&us.User.Username,
		&us.User.Email,
		&us.User.Password,
		&us.User.Role,
		&us.User.Active,
		&us.User.TenantId,
	)

	//fmt.Sprintln("User %v", us.User)
	if err != nil {
		return models.User{}, err
	}

	return us.User, nil
}

func (us *UserService) List(createdBy int) ([]models.User, error) {
	query := "SELECT * FROM users WHERE active = ?;"
	
	rows, err := us.Db.Queryx(query, createdBy)
	if err != nil {
		return []models.User{}, err 
	}

	users := []models.User{}
	for rows.Next() {
		var u models.User
		err = rows.StructScan(&u)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, u)
	}

	return users, nil

}

func (us *UserService) GetById(userId int) (models.User, error) {
	// query := "SELECT * FROM users WHERE id = ?;"
	
	// rows, err := us.Db.Queryx(query, createdBy)
	// if err != nil {
	// 	return User{}, err
	// }

	user := models.User{}
	// this will pull the first place directly into p
	err := us.Db.Get(&user, "SELECT * FROM users WHERE id = ?", userId)

    if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (us *UserService) List2() ([]models.User, error) {

	users := []models.User{}
	rows, err := us.Db.Queryx("SELECT * FROM users")
	if err != nil {
		return []models.User{}, err 
	}
	for rows.Next() {
		var p models.User
		err = rows.StructScan(&p)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, p)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

    // if users == nil {
	// 	return c.JSON(http.StatusBadRequest, "No Records Found")
	// } 
		
	// return c.JSON(http.StatusOK, users)


	return users, nil
}