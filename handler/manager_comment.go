package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) AddManagerReply(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")

	managerComment := new(model.ManagerComment)

	//here, we have restaurant id and food id !
	//the costumer had to buy this food !
	if err := c.Bind(&managerComment); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}


	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "manager not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	if res.ID != managerComment.RestaurantID {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "Sorry, You can't reply this comment!", false))
	}

	//if !foodExists(res.Foods, managerComment.FoodID) {
	//	return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "Sorry, This food doesn't exist!", false))
	//}

	parent, err := h.commentsStore.GetCommentByID(managerComment.ParentID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "sorry, The parent doesn't exist.", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	var reply model.ManagerReply
	reply.Text = managerComment.Text
	parent.Reply = reply

	err = h.commentsStore.AddManagerReply(reply, managerComment.ParentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(parent, "", true))
}

func foodExists(foods []primitive.ObjectID, foodID primitive.ObjectID) bool {
	for _, food := range foods {
		if food == foodID {
			return true
		}
	}
	return false
}