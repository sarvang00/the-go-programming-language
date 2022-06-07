package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "JTHJ",
			Year:   2014,
			Color:  true,
			Actors: []string{"SRK", "AS"},
		},
		{
			Title:  "D2",
			Year:   2010,
			Color:  true,
			Actors: []string{"HR", "AR"},
		},
		{
			Title:  "TD",
			Year:   1942,
			Color:  false,
			Actors: []string{"CC"},
		},
	}

	// Converting a Go data structure like movies to JSON is called marshaling
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(movies)

	// Marshal does not keep empty spaces. To better read it, we need an indented Marshal
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
