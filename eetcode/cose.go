package eetcode

import (
	"fmt"
	"strings"
)

func Days(days []map[string]string) []map[string]string {
	result := make([]map[string]string, 0, len(days))
	result = append(result, map[string]string{
		"days": days[0]["day"],
		"open": days[0]["open"],
		"close": days[0]["close"],
	})
	for _, day := range days[1:] {
		for i, res := range result {
			if (res["open"] == day["open"]) && (res["close"] == day["close"]) {
				comp := strings.Contains(res["days"], day["day"])
				if (!comp) {
					begin, _, _ := strings.Cut(res["days"], "-")
					res["days"] = fmt.Sprintf("%v-%v", begin, day["day"])
					break
				}
			} else if (i == len(result) - 1) {
				result = append(result, map[string]string{
					"days": day["day"],
					"open": day["open"],
					"close": day["close"],
				})
			}
		}
	}
	return result
}
