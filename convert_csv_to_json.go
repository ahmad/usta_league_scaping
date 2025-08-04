package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

type Section struct {
	SectionID    string `json:"SectionId"`
	SectionName  string `json:"SectionName"`
	DistrictID   string `json:"DistrictId"`
	DistrictName string `json:"DistrictName"`
	AreaID       string `json:"AreaId"`
	AreaName     string `json:"AreaName"`
	Link         string `json:"Link"`
}

func readCsvFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csv := csv.NewReader(f)
	records, err := csv.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func main() {
	lines := readCsvFile("USTA Leagues Section, District, and Area IDs.csv")

	sections := make([]Section, len(lines)-1)
	for i, line := range lines {
		if i == 0 {
			continue
		}

		sections[i-1] = Section{
			SectionID:    line[0],
			SectionName:  line[1],
			DistrictID:   line[2],
			DistrictName: line[3],
			AreaID:       line[4],
			AreaName:     line[5],
			Link:         line[6],
		}
	}

	j, _ := json.MarshalIndent(sections, "", "  ")
	os.WriteFile("sections.json", j, os.ModePerm)

}
