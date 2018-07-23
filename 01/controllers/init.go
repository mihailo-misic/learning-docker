package controllers

import "github.com/jinzhu/gorm"

var db *gorm.DB

func ConnectControllers(d *gorm.DB) {
	db = d
}

