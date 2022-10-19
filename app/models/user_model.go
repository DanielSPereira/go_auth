package models

import (
	"auth/pkg/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primary_key" db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	u.CreatedAt = time.Now()

	u.Password, err = utils.EncryptPassword(u.Password)
	if err != nil {
		err = errors.New("can't save invalid data")
	}

	return
}
