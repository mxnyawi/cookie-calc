package calculator

import (
	"github.com/mxnyawi/cache-calc/logger"
)

// cookie represents a cookie data
type cookie struct {
	value string
	date  string
}

// extractCookies extracts the cookies from a CSV row
func extractCookies(logger logger.LoggerInterface, data [][]string) ([]*cookie, error) {
	if len(data) == 0 {
		return nil, ErrInvalidData
	}

	cookies := make([]*cookie, len(data))
	for i, row := range data {
		if len(row) >= 2 {
			if len(row[0]) != 16 {
				logger.Error("Invalid cookie data: %v", row[0])
				continue
			}
			cookies[i] = &cookie{value: row[0], date: row[1]}
		}
	}

	if len(cookies) == 0 || cookies[0] == nil && len(cookies) == 1 {
		return nil, ErrInvalidData
	}

	return cookies, nil
}

// extractDay extracts the day from a date string
func extractDay(date string) string {
	if len(date) < 10 {
		return ""
	}

	return date[:10]
}

// Calculate calculates the most active cookies on a given date
func Calculate(logger logger.LoggerInterface, data [][]string, date string) ([]string, error) {
	logger.Info("Calculating most active cookies...")

	logger.Info("Extracting cookies from the CSV data...")
	cookies, err := extractCookies(logger, data)
	if err != nil {
		logger.Error("Error extracting cookies: %v", err)
		return nil, err
	}

	logger.Info("Counting the occurrences of each cookie on the given date...")
	counts := make(map[string]int)
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		if extractDay(cookie.date) == date {
			logger.Info("Cookie %s found on %s", cookie.value, date)
			counts[cookie.value]++
		}
	}

	logger.Info("Finding the cookies with the most occurrences...")
	var maxCount int
	maxCookies := make([]string, 0)
	for cookie, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCookies = []string{cookie}
		} else if count == maxCount {
			maxCookies = append(maxCookies, cookie)
		}
	}

	if len(maxCookies) == 0 {
		logger.Error("No cookies found for %s", date)
		return nil, ErrNoCookies
	}

	logger.Info("Most active cookies found: %v", maxCookies)
	return maxCookies, nil
}
