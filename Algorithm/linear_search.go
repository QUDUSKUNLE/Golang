package algo

//  Golang implements Linear Search
func LinearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func ReverseTwo(str string) string {
	reverse := []rune(str)
	for index, j := 0, len(reverse) - 1; index < len(reverse) / 2; index, j = index + 1, j -1 {
		reverse[index], reverse[j] = reverse[j], reverse[index]
	}
	return string(reverse)
}
