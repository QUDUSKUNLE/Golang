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

func (sess Season) String() string {
	switch sess {
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
