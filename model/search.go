package model

//empty string shows that field is not in the query
//0 value for int(Area) also shows the field is not in the query
type Search struct {
	FoodName       string `json:"food_name" bson:"food_name"`
	RestaurantName string `json:"restaurant_name" bson:"restaurant_name"`
	Area           int    `json:"area" bson:"area"`
}
