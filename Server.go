package main

import (
	_ "github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"myapp/db"
	"myapp/handler"
	"myapp/router"
	"myapp/store"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// default Port
		port = "8080"
	}

	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	mongoClient, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	usersDb := db.SetupUsersDb(mongoClient)
	restaurantsDb := db.SetupRestaurantsDb(mongoClient)
	foodsDb := db.SetupFoodsDb(mongoClient)
	ordersDb := db.SetupOrdersDb(mongoClient)
	commentsDb := db.SetupCommentsDb(mongoClient)
	ManagerCommentsDb := db.SetupManagerCommentsDb(mongoClient)
	g := r.Group("")
	userStore := store.NewUserStore(usersDb)
	restaurantStore := store.NewRestaurantStore(restaurantsDb)
	foodStore := store.NewFoodStore(foodsDb)
	orderStore := store.NewOrderStore(ordersDb)
	commentStore := store.NewCommentStore(commentsDb)
	managerCommentStore := store.NewManagerCommentStore(ManagerCommentsDb)
	h := handler.NewHandler(userStore, restaurantStore, foodStore, orderStore, commentStore, managerCommentStore)
	h.RegisterRoutes(g)

	// Fire up the trends beforehand
	//err = hs.Update()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// RUN THIS IF YOUR HASHTAG DATABASE IS EMPTY
	// StartUpTrends(ts, h)

	r.Logger.Fatal(r.Start("0.0.0.0:" + port))
	//e := echo.New()
	//e.POST("restaurant", h.CreateRestaurant)
	//
	//fmt.Println(h)
	//e.Logger.Fatal(e.Start(":1373"))

}