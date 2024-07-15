package services

import (
	"net/http"

	"github.com/ImArnav19/stocks/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := app.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (app *App) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := app.DB.Preload("Wallet").Preload("Investments").Preload("Sells").Preload("Watchlist").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (app *App) GetUserByName(c *gin.Context) {
	user_name := c.Param("username")
	var user models.User
	if err := app.DB.Preload("Wallet").Preload("Investments").Preload("Sells").Preload("Watchlist").Where("user_name = ?", user_name).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

func (app *App) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := app.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := app.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (app *App) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := app.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Permanently delete the user
	if err := app.DB.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
