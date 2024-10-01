package algo

/**
	A binary search is a search strategy used to find elements within a list
	by consistently reducing the amount of data to be searched and thereby
	increasing the rate at which the search term is found. To use a binary search algorithm,
	the list to be operated on must have already been sorted
*/
func BinarySearch(dataList []int, needle int) bool {
	low := 0;
	high := len(dataList) - 1;

	for low <= high {
		median := (low + high) / 2
		if dataList[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(dataList) || dataList[low] != needle {
		return false
	}
	return true
}
