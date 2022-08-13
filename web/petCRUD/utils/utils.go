package utils

import (
	"database/sql"
	"log"
	"time"
)

const LayoutDate string = "2006-01-02"

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertTimeToDate(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format(LayoutDate)
	}
	return ""
}

func ConvertDateToTime(d string) time.Time {
	t, err := time.Parse(LayoutDate, d)

	if err != nil {
		return time.Time{}
	}
	return t
}
