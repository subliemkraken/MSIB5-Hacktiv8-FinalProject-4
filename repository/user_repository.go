package repository

import (
	"FinalProject4/model/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	FindUserByID(ID int) (entity.User, error)
	UpdateUser(ID int, user entity.User) (entity.User, error)
	DeleteUser(ID int) (entity.User, error)
	UserTopup(ID int, user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByID(id int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ID int, user entity.User) (entity.User, error) {
	fmt.Println("dari repository ", user)
	err := r.db.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ID int) (entity.User, error) {
	userDeleted := entity.User{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&userDeleted).Error

	if err != nil {
		return entity.User{}, err
	}

	return userDeleted, nil
}

func (r *userRepository) UserTopup(ID int, user entity.User) (entity.User, error) {
	err := r.db.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
