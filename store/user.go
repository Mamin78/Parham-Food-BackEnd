package store

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type UserStore struct {
	db *mongo.Collection
}

func NewUserStore(db *mongo.Collection) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) CreateUser(user *model.User) error {
	_, err := us.db.InsertOne(context.TODO(), user)
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

func (us *UserStore) UpdateUserCredit(userID primitive.ObjectID, newCredit float64) error {
	newUser := bson.M{"credit": newCredit}
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": newUser})
	return err
}

func (us *UserStore) AddOrderToUserByID(newOrder *model.Order, user *model.User) error {
	user.Orders = append(user.Orders, newOrder.ID)
	newRes := bson.M{"orders": user.Orders}
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": newRes})
	return err
}

func (us *UserStore) AddCommentToUser(commentId primitive.ObjectID, user *model.User) error {
	user.Comments = append(user.Comments, commentId)
	newRes := bson.M{"comments": user.Comments}
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": newRes})
	return err
}
