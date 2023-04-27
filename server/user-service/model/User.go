package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Country   string `json:"country" bson:"country"`
	Street    string `json:"street" bson:"street"`
	StreetNum string `json:"streetNum" bson:"street_num"`
	ZIPCode   int    `json:"zip" bson:"zip"`
}

type Model struct {
	Id         *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedOn  int                 `json:"createdOn" bson:"created_on"`
	ModifiedOn int                 `json:"modifiedOn" bson:"modified_on"`
	DeletedOn  int                 `json:"deletedOn" bson:"deleted_on"`
}

const (
	UserRole  string = "USER"
	AdminRole string = "ADMIN"
)

type User struct {
	Model     `bson:",inline"`
	FirstName string `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string `json:"lastName" bson:"last_name" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required,email"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
	Address
}
