package main

import (
	"os"

	"github.com/billzayy/elastic-golang/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	cors.Default()
	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := app.Group("/api")
	{
		api.GET("/location/search", controller.SearchElasticController)
		api.GET("/location/", controller.AllLocationController)
	}
	app.GET("/connectDB", controller.PostgresControllerConnect)

	app.Run(":" + os.Getenv("PORT"))
}
