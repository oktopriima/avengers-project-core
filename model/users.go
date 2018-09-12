package model

import (
	"time"

	"github.com/oktopriima/go-package/library"
)

type Users struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Password      string             `json:"-"`
	RememberToken library.NullString `json:"remember_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
