// Package models
package models

type User struct {
	ID int
	UserName string
	DisplayedName string
	Password string
	CreatedAt string
}

type UserDTO struct {
	ID int `json:"ID"`
	UserName string `json:"user_name"`
	DisplayedName string `json:"displayed_name"`
	CreatedAt string `json:"created_at"`
}

func NewUserDTO(u User) UserDTO {
	return UserDTO{UserName: u.UserName}
}
