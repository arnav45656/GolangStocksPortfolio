package main

import (
	"fmt"
	"log"

	"github.com/ImArnav19/stocks/config"
	"github.com/ImArnav19/stocks/controllers"
	"github.com/ImArnav19/stocks/models"
	"github.com/ImArnav19/stocks/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Envs.DBUser, config.Envs.DBPasswd, config.Envs.DBAddr, config.Envs.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}

	db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Investment{}, &models.Sell{}, &models.Watchlist{})

	app := &services.App{DB: db}

	router := controllers.NewRouter(app)

	router.Run(":" + config.Envs.Port)

}
