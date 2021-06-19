package main

import (
	"github.com/Mamin78/Parham-Food-BackEnd/db"
	"github.com/Mamin78/Parham-Food-BackEnd/handler"
	"github.com/Mamin78/Parham-Food-BackEnd/router"
	"github.com/Mamin78/Parham-Food-BackEnd/store"
	_ "github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// default Port
		port = "8000"
	}

	r := router.New()
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
	g := r.Group("/api")
	userStore := store.NewUserStore(usersDb)
	restaurantStore := store.NewRestaurantStore(restaurantsDb)
	foodStore := store.NewFoodStore(foodsDb)
	orderStore := store.NewOrderStore(ordersDb)
	commentStore := store.NewCommentStore(commentsDb)
	managerCommentStore := store.NewManagerCommentStore(ManagerCommentsDb)
	h := handler.NewHandler(userStore, restaurantStore, foodStore, orderStore, commentStore, managerCommentStore)
	h.RegisterRoutes(g)

	r.Logger.Fatal(r.Start("0.0.0.0:" + port))

	//form := "3:04"
	//t2, _ := time.Parse(form, "8:41")
	//fmt.Println(t2)
}