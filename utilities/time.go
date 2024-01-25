package utilities

import (
	"strconv"
	"time"
)

func IsTokenExpired(expiry string, updated time.Time) (bool, error) {

	now := time.Now()
	n, err := strconv.Atoi(expiry)
	if err != nil {
		return false, err
	}

	expired := updated.Add(time.Duration(n) * time.Second)

	if now.Before(expired) {
		return true, err
	}

	return false, err
}
