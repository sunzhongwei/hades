package hades

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Time = *s
	} else {
		v.Valid = false
	}
	return nil
}

// Datetime to string, format: 2006-01-02
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// Format Datetime As English Locale, E.g: January 02, 2006
func FormatDateEn(t time.Time) string {
	return t.Format("January 02, 2006")
}

// Format Datetime As Chinese Locale, E.g: 2006年01月02日
func FormatDateCn(t time.Time) string {
	return t.Format("2006年01月02日")
}
