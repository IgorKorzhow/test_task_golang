package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"test_task_golang/configs"
	_ "test_task_golang/docs"
	"test_task_golang/internal/controllers"
	"test_task_golang/internal/database"
	"test_task_golang/internal/repositories"
	"test_task_golang/internal/services"
)

// @title		Currency courses api
// @version	1.0
// @BasePath	/api/v1
// @schemes	http
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
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/currency_courses", currencyCourseController.GetCurrencies)
	}

	r.GET("/ping", controllers.Ping)

	serverUrl := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)

	//Swagger doc
	url := ginSwagger.URL(serverUrl + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run server
	if err := r.Run(serverUrl); err != nil {
		log.Fatal("Server run failed:", err)
	}
}
