package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentStore interface {
	CreateComment(comment *model.Comment) error
	GetAllFoodComments(foodID string) (*[]model.Comment, error)
	GetCommentByID(ID primitive.ObjectID) (*model.Comment, error)
	AddManagerReply(managerReply model.ManagerReply, commentId primitive.ObjectID) error
}
