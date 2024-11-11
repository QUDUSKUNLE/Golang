package eetcode

import (
	"fmt"
	"strings"
)

func GroupOpeningDays(openingHours []map[string]string) []map[string]string {
	groups := make(map[string][]string)
	result := make([]map[string]string, 0)

	// Iterate through the days in openingHours map
	for _, day := range openingHours {
		hours := fmt.Sprintf("%s-%s", day["open"], day["close"])
		groups[hours] = append(groups[hours], day["day"])
	}

	for hours, day := range groups {
		hour := strings.Split(hours, "-")
		if len(day) > 1 {
			result = append(result, map[string]string{
				"days": fmt.Sprintf("%s-%s", day[0], day[len(day) -1]),
				"open": hour[0],
				"close": hour[len(hour)-1],
			})
		} else {
			result = append(result, map[string]string{
				"days": day[0],
				"open": hour[0],
				"close": hour[len(hour)-1],
			})
		}
	}
	return result
}
