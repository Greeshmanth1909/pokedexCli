package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
    "github.com/Greeshmanth1909/pokedexCli/pokecache"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string  `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationLog struct {
    Next string
    Previous string
}


func CommandMap() error {
     commandMapHelper(&location)
     return nil
}

func CommandMapb() error {
    commandMapbHelper(&location)
    return nil 
}

var duration time.Duration = 5 * time.Minute

// Initialise cache
var cache = pokecache.NewCache(duration)

// A location log struct that tracks urls
var location = LocationLog{
    Next: "https://pokeapi.co/api/v2/location-area/",
    Previous: "",
}

// A helper for the CommandMap function to call the api and keep track of the urls visited
func commandMapHelper(p *LocationLog) {
    next := p.Next
    resp, err := http.Get(next)

    if err != nil {
        fmt.Println("Error getting response from url")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Error in reading response body")
        return 
    }

    // Convert byte slice body to json
    var location Location
    json.Unmarshal(body, &location)

    p.Next = location.Next
    p.Previous = next

    results := location.Results
    for _, value := range results {
        fmt.Println(value.Name)
    }
    return
}

// Helper for command map back function
func commandMapbHelper(p *LocationLog) {
    previous := p.Previous

    if previous == "" {
        fmt.Println("AT THE STARTING POINT!!!, use map command to explore")
        return
    }
    resp, err := http.Get(previous)

    if err != nil {
        fmt.Println("Error getting response from url")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Error in reading response body")
        return 
    }

    // Convert byte slice body to json
    var location Location
    json.Unmarshal(body, &location)

    p.Next = location.Next
    p.Previous = location.Previous

    results := location.Results
    for _, value := range results {
        fmt.Println(value.Name)
    }
    return
}

