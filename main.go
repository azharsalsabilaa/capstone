package main

import (
	"log"
	food "mygram/Food"
	"mygram/cafe"
	"mygram/handler"
	"mygram/hotel"
	"mygram/souvenir"
	"mygram/touris"
	"mygram/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bangkit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connestion Error")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	hotelRepository := hotel.NewRepository(db)
	hotelService := hotel.NewService(hotelRepository)
	hotelHandler := handler.NewHotelHandler(hotelService)
	tourisRepository := touris.NewRepository(db)
	tourisService := touris.NewService(tourisRepository)
	tourisHandler := handler.NewTourisHandler(tourisService)
	cafeRepository := cafe.NewRepository(db)
	cafeService := cafe.NewService(cafeRepository)
	cafeHandler := handler.NewCafeHandler(cafeService)
	souvenirRepository := souvenir.NewRepository(db)
	souvenirService := souvenir.NewService(souvenirRepository)
	souvenirHandler := handler.NewSouvenirHandler(souvenirService)
	foodsRepository := food.NewRepository(db)
	foodsService := food.NewService(foodsRepository)
	foodsHandler := handler.NewFoodHandler(foodsService)

	router := gin.Default()
	api := router.Group("/users/v1")
	apiHotel := router.Group("/hotels")
	apiTouris := router.Group("/touris")
	apiCafe := router.Group("/cafe")
	apiSouvenir := router.Group("/souvenir")
	apiFoods := router.Group("food")

	api.POST("/user", userHandler.RegisterUser)
	apiHotel.GET("/", hotelHandler.GetHotel)
	apiHotel.GET("/rating", hotelHandler.GetRate)
	apiTouris.GET("/", tourisHandler.GetTouris)
	apiTouris.GET("/rate", tourisHandler.GetRate)
	apiCafe.GET("/", cafeHandler.GetCafe)
	apiCafe.GET("/rate", cafeHandler.GetRate)
	apiSouvenir.GET("/", souvenirHandler.GetSouvenir)
	apiSouvenir.GET("/rate", souvenirHandler.GetRate)
	apiFoods.GET("/", foodsHandler.GetFoods)
	apiFoods.GET("/rate", foodsHandler.GetRate)

	router.Run(":8080")
}
