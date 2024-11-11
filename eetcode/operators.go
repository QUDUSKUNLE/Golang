package eetcode

import (
	"strings"
	"fmt"
)

func ReturnWeekends(days []map[string]string) []map[string]string {
	result := make([]map[string]string, 0, 7)
	for index := 0; index < len(days); index++ {
		if index == 0 {
			result = append(result, map[string]string{
				"days": days[index]["day"],
				"open": days[index]["open"],
				"close": days[index]["close"],
			})
		} else {
			for jindex := 0; jindex < len(result); jindex++ {
				if (result[jindex]["open"] == days[index]["open"] &&
					result[jindex]["close"] == days[index]["close"]) {
						if (!strings.Contains(result[jindex]["days"], days[index]["day"])) {
							before, _, _ := strings.Cut(result[jindex]["days"],  "-")
							result[jindex]["days"] = fmt.Sprintf("%v-%v", before, days[index]["day"])
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
