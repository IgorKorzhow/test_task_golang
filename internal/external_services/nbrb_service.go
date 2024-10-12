package external_services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"test_task_golang/configs"
	"test_task_golang/internal/types"
)

type NbrbService struct {
	baseUrl string
}

func NewNbrbService(config configs.Config) *NbrbService {
	return &NbrbService{config.NBRBServiceUrl}
}

func (ns *NbrbService) sendRequest(url string) []byte {
	response, err := http.Get(ns.baseUrl + url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func (ns *NbrbService) GetCurrencyCoursesForPeriodicity(periodicity int) []types.CurrencyCourse {
	requestUrl := fmt.Sprintf("/exrates/rates?periodicity=%d", periodicity)

	responseJsonBytes := ns.sendRequest(requestUrl)

	var currencyCourses []types.CurrencyCourse
	err := json.Unmarshal(responseJsonBytes, &currencyCourses)
	if err != nil {
		log.Fatal("Happen exception while converting json into currency courses:", err)
	}

	return currencyCourses
}
