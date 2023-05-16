package repository

import (
	"booking-app/server/user-service/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(user model.User) (*model.User, error) {
	user.CreatedOn = int(time.Now().Unix())
	user.ModifiedOn = int(time.Now().Unix())

	if _, err := GetUserByEmail(user.Email); err == nil {
		return &model.User{}, err
	}
	_, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Panic("could not save document to database! err: ", err.Error())
		return &model.User{}, err
	}
	return &user, nil
}

func GetAllUsers() []model.User {
	result := []model.User{}
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic("could not find users err: ", err.Error())
		return []model.User{}
	}
	cursor.All(context.TODO(), &result)
	return result
}

func GetUserByEmail(mail string) (*model.User, error) {
	user := model.User{}
	filter := bson.M{"email": bson.M{"$eq": mail}}

	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(id primitive.ObjectID) (*model.User, error) {
	user := model.User{}

	filter := bson.M{"_id": bson.M{"$eq": id}}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user model.User) (*model.User, error) {
	user.ModifiedOn = int(time.Now().Unix())

	objID, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err != nil {
		log.Println("could not convert hex to id")
		return &model.User{}, err
	}
	edited := bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"password":   user.Password,
		"address":    user.Address,
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": edited}
	var updatedUser model.User
	err = usersCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedUser)
	if err != nil {
		return &model.User{}, err
	}

	return &updatedUser, nil
}

func DeleteUser(id string) error {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": userId}}
	if _, err := usersCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Print("Could not delete user with hex id", id)
		return err
	}

	return nil
}
