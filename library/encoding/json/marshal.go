package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Point struct {
		X, Y int
	}

	type Line struct {
		Points  []*Point
		Quality map[string]interface{} `json:"quality,string"`
	}

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		Line   *Line
	}
	group := []*ColorGroup{
		{
			1,
			"Reds",
			[]string{"Crimson", "Red", "Ruby", "Maroon"},
			nil,
		}, {
			2,
			"Blues",
			[]string{"Golang", "Blue", "Python", "Yellow"},
			&Line{
				[]*Point{
					{10, 10},
					{20, 40},
				},
				map[string]interface{}{
					"big":    98.27,
					"medium": "5px",
					"small":  5.30,
				},
			},
		},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
