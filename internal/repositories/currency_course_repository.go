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

func (cr *CurrencyCourseRepository) InsertCurrencyCourses(currencyCourses []*types.CurrencyCourse) (int64, error) {
	sqlStr := "INSERT INTO currency_courses_history(currency_type, currency_scale, currency_name, currency_official_rate, on_date) VALUES "

	var vals []interface{}
	for _, currencyCourse := range currencyCourses {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(
			vals,
			currencyCourse.CurrencyType,
			currencyCourse.CurrencyScale,
			currencyCourse.CurrencyName,
			currencyCourse.CurrencyOfficialRate,
			currencyCourse.OnDate,
		)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, err := cr.DBConnection.Prepare(sqlStr)

	if err != nil {
		log.Fatal("Happened exception white preparing sql request ", err)
	}

	result, err := stmt.Exec(vals...)

	if err != nil {
		log.Fatal("Happened exception white executing sql request ", err)
	}

	return result.RowsAffected()
}

func convertRowsToModel(rows *sql.Rows) (*types.CurrencyCourse, error) {
	var currencyCourse types.CurrencyCourse

	if err := rows.Scan(
		&currencyCourse.ID,
		&currencyCourse.CurrencyType,
		&currencyCourse.CurrencyScale,
		&currencyCourse.CurrencyName,
		&currencyCourse.CurrencyOfficialRate,
		&currencyCourse.OnDate,
	); err != nil {
		return nil, err
	}

	return &currencyCourse, nil
}
