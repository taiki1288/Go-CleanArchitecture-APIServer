package infrastructure

import (
	"Go-CleanArchitecture-APIServer/interfaces/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	// newは指定した方のポインタ型を生成する関数。(構造体の初期化)
	sqlhandler := new(SqlHandler)
	sqlhandler.db = db 
	return sqlhandler
}

// データベースへの値の挿入するメソッド
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.db.Create(value)
}

// 与えられた条件にマッチするレコードを見つけるメソッド
func (handler *SqlHandler) Find(value interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(value, where...)
}

//データベースの値を更新し、値に主キーがない場合は挿入するメソッド 
func(handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.db.Save(value)
}

// 指定された条件に合う値を削除する。値に主キーがある場合は、主キーを条件に含む。
func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.db.Delete(value)
}
