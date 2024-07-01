package algo

import (
	"fmt"
	"math/rand"
)

func QuickSort(dataList []int) []int {
	if len(dataList) < 2 {
		return dataList
	}
	left, right := 0, len(dataList) - 1
	pivot := rand.Int() % len(dataList)

	dataList[pivot], dataList[right] = dataList[right], dataList[pivot]

	result := cap(dataList) // same as len

	fmt.Println(result)

	for key := range dataList {
		if dataList[key] < dataList[right] {
			dataList[left], dataList[key] = dataList[key], dataList[left]
			left++
		}
	}
	dataList[left], dataList[right] = dataList[right], dataList[left]
	QuickSort(dataList[:left])
	QuickSort(dataList[left+1:])
	return dataList
}
