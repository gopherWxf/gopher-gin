package goft

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	dsn := "root:123456@tcp(127.0.0.1:55001)/asyncflow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &GormAdapter{DB: db}
}
