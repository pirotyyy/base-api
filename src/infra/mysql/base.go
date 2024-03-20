package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	*gorm.DB
}

func NewConnection() (*Connection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Connection{
		db,
	}, nil
}
