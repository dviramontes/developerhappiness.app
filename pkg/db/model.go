package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Active   bool   `json:"active"`
	IsBot    bool   `json:"isBot"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
	ImgUrl   string `json:"imgUrl"`
	IsAdmin  bool   `json:"isAdmin"`
	IsOwner  bool   `json:"isOwner"`
	IsNew    bool   `gorm:"-;default:false"` // virtual field, not persisted. Used for upserting
}

func (db *DB) CreateUser(u *User) error {
	if err := db.Conn.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) GetUsers() ([]User, error) {
	var users []User

	if err := db.Conn.Order("created_at asc").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
