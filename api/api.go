package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
    "log"
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


func CommandMap(str ...string) error {
     commandMapHelper(&location)
     return nil
}

func CommandMapb(str ...string) error {
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
    fmt.Println(next)

    // Check if required entry exixts in cache
    val, ok := cache.Get(next)
    if ok {
        fmt.Println(string(val))
        return 
    }

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
    strarr := ""
    for _, value := range results {
        fmt.Println(value.Name)
        strarr += fmt.Sprintf("%v\n", value.Name)
    }

    // Add to cache
    cache.Add(next, []byte(strarr))
    return
}

// Helper for command map back function
func commandMapbHelper(p *LocationLog) {
    previous := p.Previous

    // Get entry from cache, if it exists
    val, ok := cache.Get(previous)
    if ok {
        fmt.Println(string(val))
        return
    }

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
    strarr := ""
    for _, value := range results {
        fmt.Println(value.Name)
        strarr += fmt.Sprintf("%v\n", value.Name)
    }

    // Cache new entry
    cache.Add(previous, []byte(strarr))

    return
}

// Struct to store incomming encounters data
type explore struct {
    EncounterMethodRates []struct {
        EncounterMethod struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"encounter_method"`
        VersionDetails []struct {
            Rate    int `json:"rate"`
            Version struct {
                Name string `json:"name"`
                URL  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"encounter_method_rates"`
    GameIndex int `json:"game_index"`
    ID        int `json:"id"`
    Location  struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"location"`
    Name  string `json:"name"`
    Names []struct {
        Language struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"language"`
        Name string `json:"name"`
    } `json:"names"`
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
        VersionDetails []struct {
            EncounterDetails []struct {
                Chance          int           `json:"chance"`
                ConditionValues []interface{} `json:"condition_values"`
                MaxLevel        int           `json:"max_level"`
                Method          struct {
                    Name string `json:"name"`
                    URL  string `json:"url"`
                } `json:"method"`
                MinLevel int `json:"min_level"`
            } `json:"encounter_details"`
            MaxChance int `json:"max_chance"`
            Version   struct {
                Name string `json:"name"`
                URL  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"pokemon_encounters"`
}

// Explore function all the pokemon present in a given area
func Explore(areas ...string) error {
    if len(areas) != 1 {
        fmt.Println("Please specify area to explore")
        return nil
    }
    
    baseURL := "https://pokeapi.co/api/v2/location-area/"

    // get area and send request to the api
    area := areas[0]
    currentURL := fmt.Sprintf("%v%v/", baseURL, area)
    resp, err := http.Get(currentURL)

    if err != nil {
        log.Print("Error while processing explore request to the api, try again")
        return nil
    }

    // Access json body and unmarshall it
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    if err != nil {
        fmt.Println("Unable to parse json")
        return nil
    }

    var explorestr explore
    json.Unmarshal(body, &explorestr)

    // explorestr contains parsed data, extract available pokemon in the area
    for _, str := range explorestr.PokemonEncounters {
        fmt.Println(str.Pokemon.Name)
    }
    return nil
}
