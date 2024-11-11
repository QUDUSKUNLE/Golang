package eetcode

import (
	"fmt"
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
		"close": "5:00 PM",
	},
	{
		"day": "Wednesday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Thursday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Friday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Saturday",
		"open": "9:00 AM",
		"close": "2:00 PM",
	},
	{
		"day": "Sunday",
		"open": "9:00 AM",
		"close": "2:00 PM",
	},
}

var Data2 = []map[string]string{
	{
		"day": "Monday",
		"open": "8:00 AM",
		"close": "5:00 PM",
	},
	{
		"day": "Tuesday",
		"open": "8:00 AM",
		"close": "5:00 PM",
	},
	{
		"day": "Wednesday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Thursday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Friday",
		"open": "8:00 AM",
		"close": "4:00 PM",
	},
	{
		"day": "Saturday",
		"open": "9:00 AM",
		"close": "2:00 PM",
	},
}

func ReturnDays(days []map[string]string) []map[string]string {
	result := make([]map[string]string, 0, 7)
	lens := len(days)
	for index := 0; index < lens; index++ {
		if (index == 0) {
			result = append(result, map[string]string{
				"days": days[index]["day"],
				"open": days[index]["open"],
				"close": days[index]["close"],
			})
		} else {
			for jindex := 0; jindex < len(result); jindex++ {
				compare := strings.Contains(result[jindex]["days"], days[index]["day"])
				if ((days[index]["open"] == result[jindex]["open"] && days[index]["close"] == result[jindex]["close"])) {
						if (!compare) {
							begin, _, _ := strings.Cut(result[jindex]["days"], "-")
							result[jindex]["days"] = fmt.Sprintf("%v-%v", begin, days[index]["day"])
							break
						}
				} else if (jindex == len(result) - 1) {
					result = append(result, map[string]string{
						"days": days[index]["day"],
						"open": days[index]["open"],
						"close": days[index]["close"],
					})
				}
			}
		}
	}
	return result
}
