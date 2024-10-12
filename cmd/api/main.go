package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"test_task_golang/configs"
	"test_task_golang/internal/controllers"
	"test_task_golang/internal/database"
	"test_task_golang/internal/repositories"
	"test_task_golang/internal/services"
)

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		log.Fatal("Problems with configuration: ", err)
	}

	// DB initialization
	db, err := database.Connect(config)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Could not close connection to the database:", err)
		}
	}(db)

	// Gin initialization
	r := gin.Default()

	// Repositories initialization
	currencyCourseRepository := repositories.NewCurrencyCourseRepository(db)

	// Services initialization
	currencyService := services.NewCurrencyCourseService(currencyCourseRepository)

	// Controllers initialization
	currencyCourseController := controllers.NewCurrencyController(currencyService)

	// Routes
	r.GET("/currency_courses", currencyCourseController.GetCurrencies)
	// Run server
	serverUrl := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	if err := r.Run(serverUrl); err != nil {
		log.Fatal("Server run failed:", err)
	}
}
