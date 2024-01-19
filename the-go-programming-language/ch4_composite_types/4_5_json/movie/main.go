// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie
//
// omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

// 在终端执行：
//
//	go run ./ch4_composite_types/4_5_json/movie/main.go
func main() {
	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		// [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
	}

	{
		data, err := json.MarshalIndent(movies, "", "  ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//[
		//  {
		//    "Title": "Casablanca",
		//    "released": 1942,
		//    "Actors": [
		//      "Humphrey Bogart",
		//      "Ingrid Bergman"
		//    ]
		//  },
		//  {
		//    "Title": "Cool Hand Luke",
		//    "released": 1967,
		//    "color": true,
		//    "Actors": [
		//      "Paul Newman"
		//    ]
		//  },
		//  {
		//    "Title": "Bullitt",
		//    "released": 1968,
		//    "color": true,
		//    "Actors": [
		//      "Steve McQueen",
		//      "Jacqueline Bisset"
		//    ]
		//  }
		//]

		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
		// [{Casablanca} {Cool Hand Luke} {Bullitt}]
	}
}
