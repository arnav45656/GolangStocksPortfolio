package controllers

import (
	"github.com/ImArnav19/stocks/services"
	"github.com/ImArnav19/stocks/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *services.App) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(utils.RateLimitHandler())

	//Public routes
	router.POST("/login", app.Login)
	router.POST("/register", app.Register)
	router.GET("/users/username/:username", app.GetUserByName)

	//Private routes

	authorized := router.Group("/")
	authorized.Use(utils.JWTAuth())
	{
		authorized.GET("/users/:id", app.GetUser)
		authorized.POST("/users", app.CreateUser)
		authorized.PUT("/users/:id", app.UpdateUser)
		authorized.DELETE("/users/:id", app.DeleteUser)

		authorized.GET("/gemini", app.GetGeminiResponse)
		authorized.POST("/gemini")
	}

	return router
}
