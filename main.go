package main

import (
	"awesomeProject/module/components/appctx"
	"awesomeProject/module/middleware"
	"awesomeProject/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	db := db.Debug()
	appContext := appctx.NewAppContext(db)
	router := gin.Default()
	router.Use(middleware.Recover(appContext))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restaurant := router.Group("/v1")
	restaurant.POST("/restaurant", ginrestaurant.CreateRestaurant(appContext))
	restaurant.DELETE("/restaurant", ginrestaurant.DeleteRestaurant(appContext))
	restaurant.Group("", ginrestaurant.ListRestaurant(appContext))
	router.Run() // listen and serve on 0.0.0.0:8080
}
