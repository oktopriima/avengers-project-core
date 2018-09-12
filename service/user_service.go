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

	query := "SELECT id, name, email FROM users WHERE email=?"
	row := this.db.QueryRow(query, email)
	err := row.Scan(&m.ID, &m.Name, &m.Email)

	if err != nil {
		return err, nil
	}

	return nil, m
}
