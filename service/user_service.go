package service

import (
	"database/sql"

	"github.com/oktopriima/avengers-project-core/model"
	"github.com/oktopriima/avengers-project-core/repository"
)

type userService struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userService{db}
}

func (this *userService) FindByEmail(email string) (error, *model.Users) {
	m := new(model.Users)

	query := "SELECT id, name, email, password FROM users WHERE email=?"
	row := this.db.QueryRow(query, email)
	err := row.Scan(&m.ID, &m.Name, &m.Email, &m.Password)

	if err != nil {
		return err, nil
	}

	return nil, m
}

func (this userService) FindByID(ID int) (error, *model.Users) {
	m := new(model.Users)

	query := "SELECT * FROM users WHERE id = ?"
	row := this.db.QueryRow(query, ID)

	err := row.Scan(&m.ID, &m.Name, &m.Email, &m.Password, &m.RememberToken, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err, nil
	}

	return nil, m
}
