package controllers

import (
	"github.com/ImArnav19/stocks/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *services.App) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/users/:id", app.GetUser)
	router.POST("/users", app.CreateUser)
	router.PUT("/users/:id", app.UpdateUser)
	router.DELETE("/users/:id", app.DeleteUser)
	router.GET("/users/username/:username", app.GetUserByName)

	return router
}
