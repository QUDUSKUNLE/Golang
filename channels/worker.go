package channels

import (
	"fmt"
	"os"
	"time"
	"encoding/csv"
)

type city struct {
	name string
	location string
}

func createCity(record city) {
	time.Sleep(10 * time.Millisecond)
}

func ReadData(cityChan chan []city) {
	var cities []city
	csvFile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvFiles, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvFiles {
		cities = append(cities, city{
			name: line[0],
			location: line[1],
		})
	}
	cityChan <- cities
}

func Workers(cityChan chan city) {
	for val := range cityChan {
		createCity(val)
	}
}
