package models

import "gorm.io/gorm"

// swagger:model User
type UserBasic struct {
	Name     string `gorm:"not null" json:"name" binding:"min=3,max=20"`
	Email    string `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8,max=20"`
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

func (u *User) FindOne(id uint) (*User, error) {
	result := db.Debug().First(&u, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return u, nil
}
