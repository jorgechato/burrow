package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Connect to the db
func Connect(db Config) (*gorm.DB, error) {
	return gorm.Open(
		db.Dialect,
		db.toString(),
	)
}

// NewConnection create and get the db connection
func NewConnection(env string) (*gorm.DB, error) {
	db := loadFrom(env)
	return Connect(db)
}
