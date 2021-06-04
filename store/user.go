package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"myapp/model"
)

type UserStore struct {
	db *mongo.Collection
}

func NewUserStore(db *mongo.Collection) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) CreateUser(usertaurant *model.User) error {
	_, err := us.db.InsertOne(context.TODO(), usertaurant)
	return err
}

func (us *UserStore) GetUserByPhone(phoneNumber string) (*model.User, error) {
	var u model.User
	err := us.db.FindOne(context.TODO(), bson.M{"phone_number": phoneNumber}).Decode(&u)
	return &u, err
}

func (us *UserStore) UpdateUserInfoByPhone(phoneNumber string, user *model.User) error {
	newUser := bson.M{"name": user.Name, "password": user.Password, "credit": user.Credit, "area": user.Area, "address": user.Address}
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"phone_number": phoneNumber}, bson.M{"$set": newUser})
	return err
}

func (us *UserStore) GetUserByID(userID string) (*model.User, error) {
	var u model.User
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return &u, nil
	}
	err = us.db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return &u, err
}

func (us *UserStore) UpdateUserInfoByID(userID string, user *model.User) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}
	newUser := bson.M{"name": user.Name, "password": user.Password, "credit": user.Credit, "area": user.Area, "address": user.Address}
	_, err = us.db.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": newUser})
	return err
}
