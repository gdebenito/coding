package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DayAnalysis struct {
	Date     string    `json:"date"`
	Readings []Reading `json:"readings"`
}

type Reading struct {
	Id           string         `json:"id"`
	Time         int            `json:"time"`
	Contaminants map[string]int `json:"contaminants"`
}

type ContaminantPerHour struct {
	TotalReading int
	Id           string
}

func main() {
	content, err := ioutil.ReadFile("./data.json")

	if err != nil {
		panic(err)
	}

	days := parseData(content)

	// Order data by time, containing the total of the readings
	contaminants := make(map[int][]ContaminantPerHour)

	groupDataByTime(contaminants, days)

	findDeviation(contaminants)

}

func parseData(data []byte) []DayAnalysis {
	var reading []DayAnalysis
	err := json.Unmarshal([]byte(data), &reading)
	if err != nil {
		panic(err)
	}

	return reading

}

func groupDataByTime(contaminants map[int][]ContaminantPerHour, days []DayAnalysis) {
	for _, day := range days {
		for _, reading := range day.Readings {
			total := 0
			for _, contaminant := range reading.Contaminants {
				total += contaminant
			}

			contaminants[reading.Time] = append(contaminants[reading.Time], ContaminantPerHour{
				TotalReading: total,
				Id:           reading.Id,
			})
		}
	}
}

func findDeviation(contaminants map[int][]ContaminantPerHour) {
	ratio := 1.2
	for _, contaminantsPerHour := range contaminants {
		// fmt.Printf("%d - ", id)
		total := 0
		for _, contaminantSummary := range contaminantsPerHour {
			total += contaminantSummary.TotalReading
		}
		mean := total / len(contaminantsPerHour)

		// fmt.Printf("%d \n", media)

		for _, v := range contaminantsPerHour {
			if v.TotalReading > int(float64(mean)*ratio) {
				fmt.Println("mean: ", mean)
				fmt.Println("value: ", v.TotalReading)
				fmt.Println("id: ", v.Id)
				bs, err := hex.DecodeString(v.Id)
				if err != nil {
					panic(err)
				}
				fmt.Println("decoded id: " + string(bs))
			}
		}
	}
}
