package service

import (
	"FinalProject4/model/entity"
	"FinalProject4/model/input"
	"FinalProject4/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.UserRegisterInput) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	UpdateUser(ID int, input input.UserUpdateInput) (entity.User, error)
	DeleteUser(ID int) (entity.User, error)
	UserTopup(ID int, input input.UserTopupInput) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.UserRegisterInput) (entity.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "customer",
		Balance:  0,
	}

	if user.Role != "admin" && user.Role != "customer" {
		return entity.User{}, errors.New("role must be admin or customer")
	}

	if user.Balance > 100000000 {
		return entity.User{}, errors.New("balance must be less than 100000000")
	}

	createdUser, err := s.userRepository.CreateUser(user)

	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}

func (s *userService) GetUserByEmail(email string) (entity.User, error) {
	user, err := s.userRepository.FindUserByEmail(email)

	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		return entity.User{}, errors.New("no user found on that email")
	}

	return user, nil
}

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.FindUserByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		return entity.User{}, errors.New("no user found on that ID")
	}

	return user, nil
}

func (s *userService) UpdateUser(ID int, input input.UserUpdateInput) (entity.User, error) {
	userResult, err := s.userRepository.FindUserByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("no user found on that ID")
	}

	updatedUser := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Role:     input.Role,
		Balance:  input.Balance,
	}

	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

		if err != nil {
			return entity.User{}, err
		}

		updatedUser.Password = string(passwordHash)
	}

	fmt.Println("dari service ", updatedUser)

	userUpdated, err := s.userRepository.UpdateUser(ID, updatedUser)

	if err != nil {
		return entity.User{}, err
	}

	return userUpdated, nil
}

func (s *userService) DeleteUser(ID int) (entity.User, error) {
	userResult, err := s.userRepository.FindUserByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("no user found on that ID")
	}

	userDeleted, err := s.userRepository.DeleteUser(ID)

	if err != nil {
		return entity.User{}, err
	}

	return userDeleted, nil
}

func (s *userService) UserTopup(ID int, input input.UserTopupInput) (entity.User, error) {
	userResult, err := s.userRepository.FindUserByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("no user found on that ID")
	}

	userResult.Balance = userResult.Balance + input.Balance

	userTopup, err := s.userRepository.UserTopup(ID, userResult)

	if err != nil {
		return entity.User{}, err
	}

	return userTopup, nil
}
