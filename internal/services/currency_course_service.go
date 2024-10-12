package services

import (
	"test_task_golang/internal/repositories"
	"test_task_golang/internal/types"
	"time"
)

type CurrencyCourseService struct {
	currencyRepository *repositories.CurrencyCourseRepository
}

func NewCurrencyCourseService(currencyRepository *repositories.CurrencyCourseRepository) *CurrencyCourseService {
	return &CurrencyCourseService{currencyRepository}
}

func (cs *CurrencyCourseService) GetAllCurrencies() ([]*types.CurrencyCourse, error) {
	return cs.currencyRepository.GetAllCurrencies()
}

func (cs *CurrencyCourseService) GetCurrenciesForDate(date time.Time) ([]*types.CurrencyCourse, error) {
	return cs.currencyRepository.GetCurrenciesForDate(date)
}

func (cs *CurrencyCourseService) SyncApiCurrencyWithLocal() {

}
