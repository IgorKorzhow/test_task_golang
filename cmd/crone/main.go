package main

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"test_task_golang/configs"
	"test_task_golang/internal/crone"
	"test_task_golang/internal/database"
	"test_task_golang/internal/external_services"
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

	// Creat crone instance
	c := cron.New()

	// Repositories for crone jobs
	currencyCourseRepository := repositories.NewCurrencyCourseRepository(db)

	// Services for crone jobs
	currencyCourseService := services.NewCurrencyCourseService(currencyCourseRepository)
	nbrbService := external_services.NewNbrbService(config)

	//Crone jobs
	saveCurrenciesFormAPIInLocalDBJob := crone.NewSaveCurrenciesFormAPIInLocalDBJob(currencyCourseService, nbrbService)

	// Added tasks
	_, err = c.AddFunc("@daily", saveCurrenciesFormAPIInLocalDBJob.Run)
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// Run scheduler
	c.Start()

	// Stop scheduler when stopping program
	defer c.Stop()

	// lock main thread for block complete programme
	select {}
}
