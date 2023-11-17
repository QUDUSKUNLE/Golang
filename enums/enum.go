package enums

import (
	"fmt"
)

// Use of enums in Go
type Season int64

const (
	Summer Season = iota
	Autumn
	Winter
	Spring 
)

func (season Season) String() string {
	switch season {
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	case Spring:
		return "spring"
	}
	return "unknown"
}

func PrintSeason(s Season) {
	fmt.Println("season: ", s)
}
