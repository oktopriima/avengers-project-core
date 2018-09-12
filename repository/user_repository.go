package repository

import "github.com/oktopriima/avengers-project-core/model"

type UserRepository interface {
	FindByEmail(email string) (error, *model.Users)
	FindByID(ID int) (error, *model.Users)
}
