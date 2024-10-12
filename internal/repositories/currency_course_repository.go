package repositories

import (
	"database/sql"
	"log"
	"test_task_golang/internal/types"
	"time"
)

type CurrencyCourseRepository struct {
	DBConnection *sql.DB
}

func NewCurrencyCourseRepository(db *sql.DB) *CurrencyCourseRepository {
	return &CurrencyCourseRepository{db}
}

func (cr *CurrencyCourseRepository) GetAllCurrencies() ([]*types.CurrencyCourse, error) {
	rows, err := cr.DBConnection.Query(`
		SELECT * 
		FROM currency_courses_history
	`)

	if err != nil {
		log.Fatal("Happened exception white getting currency courses", err)
		return nil, err
	}

	var currencyCourses []*types.CurrencyCourse
	for rows.Next() {
		currencyCourse, err := convertRowsToModel(rows)
		if err != nil {
			log.Fatal("Happened exception white getting currency courses", err)
		}

		currencyCourses = append(currencyCourses, currencyCourse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencyCourses, nil
}

func (cr *CurrencyCourseRepository) GetCurrenciesForDate(date time.Time) ([]*types.CurrencyCourse, error) {
	rows, err := cr.DBConnection.Query(`
		SELECT * 
		FROM currency_courses_history
		WHERE DATE(on_date) = ?
	`, date.Format("2006-01-02"))

	if err != nil {
		log.Fatal("Happened exception white getting currency courses", err)
		return nil, err
	}

	var currencyCourses []*types.CurrencyCourse
	for rows.Next() {
		currencyCourse, err := convertRowsToModel(rows)
		if err != nil {
			log.Fatal("Happened exception white getting currency courses", err)
		}

		currencyCourses = append(currencyCourses, currencyCourse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencyCourses, nil
}

func convertRowsToModel(rows *sql.Rows) (*types.CurrencyCourse, error) {
	var currencyCourse types.CurrencyCourse

	if err := rows.Scan(
		&currencyCourse.ID,
		&currencyCourse.CurrencyType,
		&currencyCourse.CurrencyScale,
		&currencyCourse.CurrencyName,
		&currencyCourse.OnDate,
	); err != nil {
		return nil, err
	}

	return &currencyCourse, nil
}
