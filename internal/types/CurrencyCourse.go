package types

import (
	"encoding/json"
	"time"
)

type CurrencyCourse struct {
	ID                   int
	CurrencyType         string
	CurrencyScale        int
	CurrencyName         string
	CurrencyOfficialRate float64
	OnDate               time.Time
}

func (cc *CurrencyCourse) UnmarshalJSON(b []byte) error {
	var temp struct {
		ID                   int
		CurrencyType         string  `json:"Cur_Abbreviation"`
		CurrencyScale        int     `json:"Cur_Scale"`
		CurrencyName         string  `json:"Cur_Name"`
		CurrencyOfficialRate float64 `json:"Cur_OfficialRate"`
		OnDate               string  `json:"Date"`
	}

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	cc.CurrencyType = temp.CurrencyType
	cc.CurrencyScale = temp.CurrencyScale
	cc.CurrencyName = temp.CurrencyName
	cc.CurrencyOfficialRate = temp.CurrencyOfficialRate
	t, err := time.Parse("2006-01-02T15:04:05", temp.OnDate)
	if err != nil {
		return err
	}
	cc.OnDate = t

	return nil
}
