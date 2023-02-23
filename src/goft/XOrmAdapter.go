package goft

import (
	"log"
	"xorm.io/xorm"
)

type XOrmAdapter struct {
	*xorm.Engine
}

func NewXOrmAdapter() *XOrmAdapter {
	dsn := "root:123456@tcp(127.0.0.1:55001)/asyncflow?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal(engine)
	}
	engine.DB().SetMaxIdleConns(5)
	engine.DB().SetMaxOpenConns(10)
	return &XOrmAdapter{Engine: engine}
}
