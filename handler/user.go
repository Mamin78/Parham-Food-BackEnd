package handler

import (
	"fmt"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
)

func (h *Handler) userSignUp(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	if user.PhoneNumber == ""{
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "phone number must not be empty", false))
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "Invalid Phone number!", false))
	}

	if err := CheckPasswordLever(user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, err.Error(), false))
	}

	user.ID = primitive.NewObjectID()
	user.Credit = 1000000
	pass, err := PasswordToHash(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "This user existed.", false))
	}
	user.Password = pass

	err = h.userStore.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "This user existed.", false))
	}
	user.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(user, "", true))
}

func (h *Handler) UserLogin(c echo.Context) error {
	fmt.Println(c)
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request1", false))
	}

	//fmt.Println(user)
	//if err := c.Validate(user); err != nil {
	//	return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "Invalid Phone number!", false))
	//}

	userInformation, err := h.userStore.GetUserByPhone(user.PhoneNumber)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
	}

	//here we should check the password!
	if !PasswordsAreSame(userInformation.Password, user.Password) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "password is incorrect.", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(newUserResponse(userInformation), "", true))
}

func (h *Handler) UpdateUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	newUser := model.NewUser(userInformation)
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.userStore.UpdateUserInfoByPhone(userPhone, newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	newUser.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(newUser, "", true))
}

func (h *Handler) GetUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	userInformation.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(userInformation, "", true))
}

func (h *Handler) GetUserFavoriteFoods(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	allFoods, err := h.foodsStore.GetAllFoods()
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "there is some problems with foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	allUserOrders, err := h.ordersStore.GetAllUserOrders(userInformation.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "there is some problems with foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	IDs := createSliceOfIDs(getAllFoodsOrderedMoreThanFiveByUser(getAllFoodsStaredMoreThanThree(allFoods, userInformation.ID), allUserOrders))
	result, err := h.foodsStore.GetAllFavoriteFoods(IDs)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "there is some problems with foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFoods(h,result), "", true))
}

func (h *Handler) GetHistoryOfUserOrders(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	allUserOrders, err := h.ordersStore.GetAllUserOrders(userInformation.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "there is some problems with foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(allUserOrders, "", true))
}

func getAllFoodsStaredMoreThanThree(allFoods *[]model.Food, userID primitive.ObjectID) map[primitive.ObjectID]bool {
	foodSet := make(map[primitive.ObjectID]bool)
	for _, food := range *allFoods {
		for _, rate := range food.Rates {
			if rate.Rate > 3 && rate.UserID == userID {
				foodSet[food.ID] = true
			}
		}
	}

	return foodSet
}

func getAllFoodsOrderedMoreThanFiveByUser(foodSet map[primitive.ObjectID]bool, orders *[]model.Order) map[primitive.ObjectID]bool {
	foodSetNumber := make(map[primitive.ObjectID]int)
	for _, order := range *orders {
		for _, food := range order.Foods {
			foodSetNumber[food.FoodID]++
		}
	}

	for k, v := range foodSetNumber {
		if v > 5 {
			foodSet[k] = true
		}
	}

	return foodSet
}

func createSliceOfIDs(set map[primitive.ObjectID]bool) []primitive.ObjectID {
	var res []primitive.ObjectID
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}
