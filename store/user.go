package store

import (
	"go.mongodb.org/mongo-driver/bson"
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

func (rs *UserStore) CreateUser(restaurant *model.User) error {
	_, err := rs.db.InsertOne(context.TODO(), restaurant)
	return err
}

func (rs *UserStore) GetUserByPhone(phoneNumber string) (*model.User, error) {
	var u model.User
	err := rs.db.FindOne(context.TODO(), bson.M{"phone_number": phoneNumber}).Decode(&u)
	return &u, err
}
