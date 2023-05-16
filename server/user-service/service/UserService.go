package service

import (
	"booking-app/server/user-service/dto"
	"booking-app/server/user-service/helper"
	"booking-app/server/user-service/model"
	"booking-app/server/user-service/repository"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignupAdmin(user model.User) (*model.User, error) {
	user.Role = model.AdminRole
	return repository.CreateUser(user)
}

func SignupGuest(user model.User) (*model.User, error) {
	user.Role = model.UserRole
	return repository.CreateUser(user)
}

func LoginUser(loginInfo dto.LoginDTO) (string, error) {
	user, err := repository.GetUserByEmail(loginInfo.Email)
	if err != nil {
		return "", err
	}
	if user.Password != loginInfo.Password {
		return "", errors.New("incorrect password")
	}
	token, err := helper.GenerateToken(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetAll() []model.User {
	return repository.GetAllUsers()
}

func GetByEmail(email string) (*model.User, error) {
	return repository.GetUserByEmail(email)
}

func GetById(id primitive.ObjectID) (*model.User, error) {
	return repository.GetUserById(id)
}

func UpdateUser(user model.User) (*model.User, error) {
	return repository.UpdateUser(user)
}

// func IncrementCancellationsCounter(userId primitive.ObjectID) (*model.User, error) {
// 	user, err := GetById(userId)
// 	if err != nil {
// 		return &model.User{}, err
// 	}
// 	user.CancellationsCounter++
// 	user, err = UpdateUser(*user)
// 	if err != nil {
// 		return &model.User{}, err
// 	}
// 	return user, nil
// }

// func verifyPassword(dbPassword string, dtoPassword string) bool {
// 	return dbPassword == dtoPassword
// }
