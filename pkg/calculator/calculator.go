package calculator

import "log"

// Cookie represents a cookie data
type Cookie struct {
	Value string
	Date  string
}

// ExtractCookies extracts the cookies from a CSV row
func ExtractCookies(data [][]string) ([]*Cookie, error) {
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

	if len(cookies) == 0 {
		return nil, ErrInvalidData
	}

	return cookies, nil
}

// ExtractDay extracts the day from a date string
func ExtractDay(date string) string {
	if len(date) < 10 {
		return ""
	}

	return date[:10]
}

// Calculate calculates the most active cookies on a given date
func Calculate(cookies []*Cookie, date string) ([]string, error) {
	// Create a map to count the occurrences of each cookie on the given date
	counts := make(map[string]int)
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		if ExtractDay(cookie.Date) == date {
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
