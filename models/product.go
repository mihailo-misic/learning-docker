package models

import (
	"time"
)

type (
	// Product describes a Product type
	Product struct {
		ID    uint64 `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
		Name  string `json:"name" binding:"required"`
		Price int    `json:"price" binding:"required"`
		Image string `json:"image"`

		CreatedAt time.Time  `json:"created-at"`
		UpdatedAt time.Time  `json:"updated-at"`
		DeletedAt *time.Time `json:"deleted-at" sql:"index"`
	}

	// ResProduct represents a formatted Product
	ResProduct struct {
		ID    uint64 `form:"id" json:"id"`
		Name  string `form:"name" json:"name"`
		Price int    `form:"price" json:"price"`
		Image string `form:"image" json:"image"`
	}

	// FormProductStruct represents Product form data
	FormProductStruct struct {
		Name  string `form:"name" json:"name" binding:"required"`
		Price int    `form:"price" json:"price" binding:"required"`
		Image string `form:"image" json:"image"`
	}
)

func (FormProductStruct) TableName() string {
	return "products"
}
