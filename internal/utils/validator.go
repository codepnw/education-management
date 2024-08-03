package utils

import (
	"errors"
	"regexp"
	"time"
)

func IsValidDate(date string) (string, error) {
	// YYYY-MM-DD
	dateFormat := "2006-01-02"
	dob, err := time.Parse(dateFormat, date)
	if err != nil {
		return "", errors.New("invalid date format: YYYY-MM-DD")
	}
	date = dob.Format(dateFormat)
	return date, nil
}

func IsValidPhone(phone string) bool {
	rex := regexp.MustCompile(`^\+?[0-9]{10}$`)
	return rex.MatchString(phone)
}
