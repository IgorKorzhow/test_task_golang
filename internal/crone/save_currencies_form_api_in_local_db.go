package crone

import (
	"test_task_golang/internal/external_services"
	"test_task_golang/internal/services"
)

type SaveCurrenciesFormAPIInLocalDBJob struct {
	currencyService *services.CurrencyCourseService
	nbrbService     *external_services.NbrbService
}

func NewSaveCurrenciesFormAPIInLocalDBJob(cs *services.CurrencyCourseService, ns *external_services.NbrbService) SaveCurrenciesFormAPIInLocalDBJob {
	return SaveCurrenciesFormAPIInLocalDBJob{cs, ns}
}

func (job *SaveCurrenciesFormAPIInLocalDBJob) Run() {
	job.currencyService.SyncApiCurrencyWithLocal(job.nbrbService, 1)
}
