package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewSQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected To SQL database ... ); ")

	return &MySQLStorage{db: db}

}
func (s *MySQLStorage) Init() (*sql.DB, error) {
	return s.db, nil

}
