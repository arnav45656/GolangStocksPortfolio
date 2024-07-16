package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ImArnav19/stocks/config"
	"github.com/ImArnav19/stocks/gemini"
	"github.com/ImArnav19/stocks/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

func (app *App) Register(c *gin.Context) {
	var user models.UserPayload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	if err := app.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (app *App) Login(c *gin.Context) {
	var user models.UserPayload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dbUser models.UserPayload
	if err := app.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  dbUser.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	fmt.Println(config.Envs.SECRET)

	tokenString, err := token.SignedString([]byte(config.Envs.SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (app *App) GetGeminiResponse(c *gin.Context) {

	resp := gemini.GetPmt1()

	c.JSON(http.StatusOK, gin.H{"response": resp})

}

func (app *App) PostGeminiResponse(c *gin.Context) {
	var req models.GeminiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := gemini.GetData(req.Prompt, c)
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
