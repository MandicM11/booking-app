package dto

import (
	"booking-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginDTO struct {
	Email    string `json:"Email" binding:"required,Email,min=4"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterUserDTO struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"Email" binding:"required,Email,min=4"`
	Password  string `json:"password" binding:"required,min=5"`
	Country   string `json:"country"`
	Street    string `json:"street"`
}

type UpdateUserDTO struct {
	Id        *primitive.ObjectID `json:"id" binding:"required"`
	FirstName string              `json:"firstName" binding:"required"`
	LastName  string              `json:"lastName" binding:"required"`
	Email     string              `json:"email" binding:"required,email"`
	Password  string              `json:"password" binding:"required,min=5"`
	Country   string              `json:"country"`
	Street    string              `json:"street"`
}

func RegisterUserDTOToUser(user RegisterUserDTO) model.User {
	return model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Address: model.Address{
			Country: user.Country,
			Street:  user.Street,
		},
	}
}

func UpdateUserDTOtoUser(userDto UpdateUserDTO) model.User {
	return model.User{
		IdModel: model.IdModel{
			Id: userDto.Id,
		},
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		Password:  userDto.Password,
		Address: model.Address{
			Country: userDto.Country,
			Street:  userDto.Street,
		},
	}
}
