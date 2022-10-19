package queries

import (
	"auth/app/models"

	"gorm.io/gorm"
)

type UserQueries struct {
	*gorm.DB
}

func (q *UserQueries) GetUserById(id int) (models.User, error) {
	user := &models.User{ID: id}

	err := q.DB.Take(user).Error

	return *user, err
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := &models.User{Email: email}

	err := q.DB.Take(user).Error

	return *user, err
}

func (q *UserQueries) CreateUser(user *models.User) error {
	err := q.DB.Save(user).Error

	return err

}
