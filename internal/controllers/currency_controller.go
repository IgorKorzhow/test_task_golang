package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test_task_golang/internal/services"
	"time"
)

type CurrencyCourseController struct {
	currencyService *services.CurrencyCourseService
}

func NewCurrencyController(service *services.CurrencyCourseService) *CurrencyCourseController {
	return &CurrencyCourseController{service}
}

func (cc *CurrencyCourseController) GetCurrencies(c *gin.Context) {
	dateStr := c.Query("date")

	if dateStr != "" {
		date, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unsupported date format, expecting(dd.mm.yy)"})
			return
		}

		currencyCourses, err := cc.currencyService.GetCurrenciesForDate(date)

		if err != nil {
			log.Println("Happened exception while getting data", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Happened exception while getting data. Plz try again later.",
			})
			return
		}

		c.JSON(http.StatusOK, currencyCourses)
		return
	}

	currencyCourses, err := cc.currencyService.GetAllCurrencies()
	if err != nil {
		log.Println("Happened exception while getting data", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Happened exception while getting data. Plz try again later.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, currencyCourses)
}
