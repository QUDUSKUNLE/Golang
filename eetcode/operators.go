package eetcode

import (
	"strings"
)

func ReturnWeekends(days []map[string]string) []map[string]string {
	if len(days) == 0 {
		return nil
	}
	result := make([]map[string]string, 0, 7)
	for i := 0; i < len(days); i++ {
		found := false
		for j := 0; j < len(result); j++ {
			if result[j]["open"] == days[i]["open"] && result[j]["close"] == days[i]["close"] {
				// Merge days if open/close times match
				if !strings.Contains(result[j]["days"], days[i]["day"]) {
					result[j]["days"] = result[j]["days"] + "," + days[i]["day"]
				}
				found = true
				break
			}
		}
		if !found {
			result = append(result, map[string]string{
				"days":  days[i]["day"],
				"open":  days[i]["open"],
				"close": days[i]["close"],
			})
		}
	}
	return result
}
