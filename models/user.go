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

		CreatedAt time.Time  `json:"created-at" sql:"DEFAULT:current_timestamp"`
		UpdatedAt time.Time  `json:"updated-at" sql:"DEFAULT:current_timestamp"`
		DeletedAt *time.Time `json:"deleted-at" sql:"index"`
	}
	// ResUser represents a formatted User
	ResUser struct {
		ID        uint64 `form:"id" json:"id"`
		Email     string `form:"email" json:"email"`
		Password  string `form:"password" json:"password"`
		FirstName string `form:"first-name" json:"first-name"`
		LastName  string `form:"last-name" json:"last-name"`
		Image     string `form:"image" json:"image"`
	}

	// FormUserStruct represents User form data
	FormUserStruct struct {
		Email     string `form:"email" json:"email" binding:"required"`
		Password  string `form:"password" json:"password" binding:"required"`
		FirstName string `form:"first-name" json:"first-name"`
		LastName  string `form:"last-name" json:"last-name"`
		Image     string `form:"image" json:"image"`
	}
)

func (FormUserStruct) TableName() string {
	return "users"
}
