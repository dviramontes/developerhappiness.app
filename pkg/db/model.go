package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	name     string
	active   bool
	isBot    bool
	email    string
	timezone string
	imgUrl   string
	isAdmin  bool
	isOwner  bool
}

func (db *DB) CreateUser(user *User) error {
	if err := db.conn.Create(user).Error; err != nil {
		return err
	}

	return nil
}
