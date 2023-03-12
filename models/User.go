package models

import "gorm.io/gorm"

// swagger:model User
type UserBasic struct {
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

type User struct {
	gorm.Model `json:"-"`
	UserBasic
}

var db, _ = DB()

func (u *User) Create() {
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *User) FindOneByEmail(email string) (*User, error) {
	result := db.Debug().Where(&UserBasic{Email: email}).First(&u)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return u, nil
}
