package crone

import "test_task_golang/internal/services"

type SaveCurrenciesFormAPIInLocalDBJob struct {
	currencyService *services.CurrencyCourseService
}

func (job *SaveCurrenciesFormAPIInLocalDBJob) run() error {
	job.currencyService
}
