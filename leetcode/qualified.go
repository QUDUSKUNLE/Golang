package leetcode

import (
	"strings"
)

var Data = []map[string]string{
	{
		"day": "Monday",
		"open": "8:00 AM",
		"close": "5:00 PM",
	},
	{
		"day": "Tuesday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Wednesday",
		"open": "8:00 AM",
		"close": "3:00 PM",
	},
	{
		"day": "Thursday",
		"open": "8:00 AM",
		"close": "3:00 PM",
	},
}

func ReturnDays(days []map[string]string) []map[string]string {
	result := make([]map[string]string, 0)
	for _, day := range days {
		if (len(result) == 0) {
			result = append(result, map[string]string{
				"days": day["day"],
				"open": day["open"],
				"close": day["close"],
			})
		} else {
			for _, d := range result {
				if (day["open"] == d["open"]) && (day["close"] == d["close"] && !strings.Contains(d["days"], day["day"])) {
					d["days"] = d["days"] + "-" + day["day"]
				} else {
					result = append(result, map[string]string{
						"days": day["day"],
						"open": day["open"],
						"close": day["close"],
					})
				}
			}
		}
	}
	return result
	
}
