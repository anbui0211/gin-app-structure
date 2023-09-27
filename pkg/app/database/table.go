package database

import "gorm.io/gorm"

// UserCol ...
func UserCol() *gorm.DB {
	return db.Table(userTable)
}

const (
	userTable = "users"
)
