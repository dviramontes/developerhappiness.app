package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type DB struct {
	Conn *gorm.DB
}

func new(db *gorm.DB) *DB {
	return &DB{db}
}

func (db *DB) Migrate() {
	db.Conn.AutoMigrate(&User{})
}

func (db *DB) Seed() error {
	owner := User{
		Name:     "david",
		Active:   true,
		IsBot:    false,
		Email:    "dviramontes@gmail.com",
		Timezone: "America/Denver",
		ImgUrl:   "https://secure.gravatar.com/avatar/fe5373af89a931ab1660970a9b25ff2c.jpg?s=32&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0010-32.png",
		IsAdmin:  true,
		IsOwner:  true,
	}

	if err := db.Conn.
		Where(&User{Email: owner.Email}).
		Attrs(User{IsNew: true}).FirstOrCreate(&owner).Error; err != nil {
		return err
	}

	return nil
}

func Connect(connStr string) (*DB, error) {
	pgdb, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	log.Println("database connection successfully opened")

	// init db models and migrate
	db := new(pgdb)
	db.Migrate()

	log.Println("database migrated")

	if err := db.Seed(); err != nil {
		return nil, fmt.Errorf("failed to seed database, %v", err)
	}

	return db, nil
}
