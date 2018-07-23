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
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Image string `json:"image"`
	}
)
