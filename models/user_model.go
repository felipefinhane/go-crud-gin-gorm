package models

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

func (u *User) TableName() string {
	return "user"
}
