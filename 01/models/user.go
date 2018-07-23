package models

import (
	"time"
)

type (
	// User describes a User type
	User struct {
		ID        uint64 `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
		Email     string `json:"email" gorm:"UNIQUE_INDEX;not null;type:varchar(100);" binding:"required"`
		Password  string `json:"password" binding:"required"`
		FirstName string `json:"first-name" gorm:"column:first_name" binding:"required"`
		LastName  string `json:"last-name" gorm:"column:last_name" binding:"required"`
		Image     string `json:"image"`

		CreatedAt time.Time  `json:"created-at"`
		UpdatedAt time.Time  `json:"updated-at"`
		DeletedAt *time.Time `json:"deleted-at" sql:"index"`
	}
	// ResUser represents a formatted User
	ResUser struct {
		ID        uint64 `json:"id"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first-name"`
		LastName  string `json:"last-name"`
		Image     string `json:"image"`
	}
)
