package model

import "time"

type User struct {
	Id           uint       `json:"id"`
	UniqueNumber int        `json:"unique_number" validate:"unique=unique_number"`
	Name         string     `json:"name"  validate:"required"`
	Username     string     `json:"username" validate:"required,min=6,unique=username"`
	Password     string     `json:"password" validate:"required,min=6"`
	LastLogin    *time.Time `json:"last_login"`
}
