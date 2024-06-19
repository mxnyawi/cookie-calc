package calculator

import (
	"log"
)

// Cookie represents a cookie data
type Cookie struct {
	Value string
	Date  string
}

// extractCookies extracts the cookies from a CSV row
func extractCookies(data [][]string) ([]*Cookie, error) {
	cookies := make([]*Cookie, len(data))
	for i, row := range data {
		if len(row) >= 2 {
			if len(row[0]) != 16 {
				log.Printf("Invalid cookie value: %v", row[0])
				continue
			}
			cookies[i] = &Cookie{Value: row[0], Date: row[1]}
		}
	}

	if cookies[0] == nil {
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
func Calculate(data [][]string, date string) ([]string, error) {
	// Extract the cookies from the CSV data
	cookies, err := extractCookies(data)
	if err != nil {
		return nil, err
	}

	// Create a map to count the occurrences of each cookie on the given date
	counts := make(map[string]int)
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		if extractDay(cookie.Date) == date {
			counts[cookie.Value]++
		}
	}

	// Find the cookies with the most occurrences
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
		return nil, ErrNoCookies
	}

	return maxCookies, nil
}
