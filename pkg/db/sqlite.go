package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	DB *sql.DB
}

func NewSQLiteDB(databasePath string) (*SQLiteDB, error) {
	_, err := os.Stat(databasePath)
	if os.IsNotExist(err) {
		// Create the database file if it doesn't exist
		file, err := os.Create(databasePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return &SQLiteDB{DB: db}, nil
}

func (s *SQLiteDB) Close() {
	if s.DB != nil {
		if err := s.DB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}
}

func (s *SQLiteDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := s.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return result, nil
}

func (s *SQLiteDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return rows, nil
}

func (s *SQLiteDB) QueryRow(query string, args ...interface{}) *sql.Row {
	row := s.DB.QueryRow(query, args...)
	return row
}
