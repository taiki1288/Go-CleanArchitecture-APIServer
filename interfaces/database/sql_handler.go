package database

import "gorm.io/gorm"

type SqlHandler interface {
	Create(interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}