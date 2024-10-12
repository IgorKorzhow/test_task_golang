package types

import "time"

type CurrencyCourse struct {
	ID            int
	CurrencyType  string
	CurrencyScale int
	CurrencyName  string
	OnDate        time.Time
}
