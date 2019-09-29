package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Season struct used for Unmarshling the JSON object.
type Season struct {
	Name   string   `json:"name"`
	Rounds []rounds `json:"rounds"`
}

// Sub struct needed for the Season struct.
type rounds struct {
	Name    string    `json:"name"`
	Matches []matches `json:"matches"`
}

// Sub struct needed for the rounds struct
type matches struct {
	Date   string `json:"date"`
	Team1  team   `json:"team1"`
	Team2  team   `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

// team struct needed for the matches struct
type team struct {
	Name string `json:"name"`
}

// Fetches data from the API for a certain season, 2016-17 for example.
func (s *Season) fetchAPI(interval string) {
	response, _ := http.Get("https://raw.githubusercontent.com/opendatajson/football.json/master/" + interval + "/en.1.json")
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &s)

}

func (s *Season) insertAll() {
	fmt.Println(s.Name)
	for _, round := range s.Rounds {
		fmt.Println(round.Name)
		for _, match := range round.Matches {
			db.insertMatch(s.Name, match.Date, match.Team1.Name, match.Team2.Name, match.Score1, match.Score2)
			fmt.Println("Data inserted succesfull.")
		}

	}
}
