package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Player struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	City       string `json:"city"`
	State      string `json:"state"`
	Rating     string `json:"rating"`
	RatingType string `json:"rating_type"`
	RatingDate string `json:"rating_date"`
	UstaId     string `json:"usta_id"`
}

type Payload struct {
	SectionId  string   `json:"sectionId"`
	DistrictId string   `json:"districtId"`
	AreaId     string   `json:"areaId"`
	Players    []Player `json:"players"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running..."))
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t Payload
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		// fmt.Println("SectionId:", SectionId)
		// fmt.Println("DistrictId:", DistrictId)
		// fmt.Println("AreaId:", AreaId)
		// fmt.Println("Payload:", Payload)

		filename := "players/" + t.SectionId + "_" + t.DistrictId + "_" + t.AreaId + ".json"
		fmt.Println("Writing to file:", filename)

		j, _ := json.MarshalIndent(t, "", "  ")
		err = os.WriteFile(filename, j, os.ModePerm)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ \"success\": true }"))
	})

	fmt.Println("Server is running http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
