package infrastructure

import "github.com/jinzhu/gorm"

type SqlHandler struct {
	db *gorm.DB
}