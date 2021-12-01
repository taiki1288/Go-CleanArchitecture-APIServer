package infrastructure

import (

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
) 

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "user:pass@tcp(127.0.0.1:3306)/Go-CleanArchitecture-APIServer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlhandler := new(SqlHandler)
	sqlhandler.db = db 
	return sqlhandler
}