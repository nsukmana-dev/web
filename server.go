package main

import (
	"time"
	"web/config"
	"web/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/home", controller.GetHome)
		v1.GET("/edit/:id", controller.GetUserById)
		v1.GET("/detail/:id", controller.GetDetailUser)
		v1.GET("/atasan", controller.GetAtasan)
		v1.GET("/atasanbyid/:id", controller.GetAtasanById)
		v1.POST("/add", controller.PostUser)
		v1.PUT("/update/:id", controller.UpdateUser)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8081"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Run()
}
