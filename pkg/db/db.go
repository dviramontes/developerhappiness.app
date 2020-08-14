package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	conn *gorm.DB
}

func new(db *gorm.DB) *DB {
	return &DB{db}
}

func (db *DB) Migrate() {
	db.conn.AutoMigrate(&User{})
}

func (db *DB) Seed() error {
	return nil
}

func Connect(connStr string) (*DB, error) {
	pgdb, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer pgdb.Close()

	// init db models and migrate
	db := new(pgdb)
	db.Migrate()

	if err := db.Seed(); err != nil {
		return nil, fmt.Errorf("failed to seed database, %v", err)
	}

	return db, nil
}
