package codewar

/*
Write a function that returns the total surface area and volume of a box as an array: [area, volume]
*/ 

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * ((w * h) + (h * d) + (d * w)), w * h * d}
}
