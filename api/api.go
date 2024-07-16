package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap() error {
    resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        fmt.Println("error fetching the request...")
        return nil
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error in reading response body")
        return nil
    }
    // Convert byte slice body to json
    var location Location
    json.Unmarshal(body, &location)

    results := location.Results
    for _, value := range results {
        fmt.Println(value.Name)
    }

    return nil

}
