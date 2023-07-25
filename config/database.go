package config

import (
	"testing-echo/domain"
	"testing-echo/helper"

	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	connInfo = "postgres://postgres:testpass@postgres:5432/postgres?sslmode=disable"
)

func DatabaseConnectionOrm() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	db.Table("person").AutoMigrate(&domain.Person{})

	return db
}

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
