package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "math/rand"
    "time"
    "encoding/json"
    "log"
    "strings"
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

    strarr = strings.TrimSuffix(strarr, "\n")

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

    strarr = strings.TrimSuffix(strarr, "\n")
    // Cache new entry
    cache.Add(previous, []byte(strarr))

    return
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

    // Check weather the required entry already exists in cache
    val, ok := cache.Get(currentURL)
    if ok {
        fmt.Println(string(val))
        return nil
    }

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
    
    cacheString := ""

    // explorestr contains parsed data, extract available pokemon in the area
    for _, str := range explorestr.PokemonEncounters {
        fmt.Println(str.Pokemon.Name)
        cacheString += fmt.Sprintf("%v\n", str.Pokemon.Name)
    }
    cacheString = strings.TrimSuffix(cacheString, "\n")
    cache.Add(currentURL, []byte(cacheString))
    return nil
}

// Global declaration of pokedex for the catch and inspect functions
var pokedex = make(map[string]Pokemon)

/* The catch function gets information about a given pokemon, determines how easy it is to catch it based
on the "base experience" stat and adds the pokemon to the user's pokedex if its caught */
func Catch(pokemon ...string) error {
    baseURL := "https://pokeapi.co/api/v2/pokemon/"
    currentURL := fmt.Sprintf("%v%v", baseURL, pokemon[0])
    
    resp, err := http.Get(currentURL)

    if err != nil {
        fmt.Println("Error while processing get request, please try again")
        return nil
    }

    defer resp.Body.Close()
    
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Unable to parse response body")
        return nil
    }

    // data is a byte array, unmarshall it
    var pokestruct Pokemon
    json.Unmarshal(data, &pokestruct)

    baseXp := pokestruct.BaseExperience
    num := rand.Intn(baseXp)
    fmt.Println("Catching", pokemon[0])

    // the pokemon will be caught if the random number generated is in the range of ten numbers in the middle of [0, baseXp)
    lowerLimit := baseXp / 2
    upperLimit := lowerLimit + 10

    // if baseXp of pokemon is less than 10, catch it on the first attempt
    if baseXp <= 20 {
        fmt.Println("caught", pokemon[0])
        addPokemon(&pokedex, pokemon[0], pokestruct)
        return nil
    }
    
    if num >= lowerLimit && num < upperLimit {
        fmt.Println("caught", pokemon[0])
        addPokemon(&pokedex, pokemon[0], pokestruct)
        return nil
    }else {
        fmt.Println(pokemon[0], "escaped!")
    }
    return nil
}

// addPokemon function adds a given pokemon to the user's pokedex
func addPokemon(p *map[string]Pokemon, name string, data Pokemon) {
    _, ok := (*p)[name]
    if ok {
        fmt.Printf("Pokemon %v has already been caught before", name)
        return
    }
    (*p)[name] = data
}

// Inspect function prints information about a given pokemon to the console
func Inspect(pokemon ...string) error{
    val, ok := pokedex[pokemon[0]]
    if !ok {
        fmt.Printf("Pokemon %v not registered in the pokedex, catch the pokemon to add it to your pokedex")
        return nil
    }
    fmt.Printf("Name: %v\n", val.Name)
    fmt.Printf("Height: %v\n", val.Height)
    fmt.Printf("Weight: %v\n", val.Weight)
    fmt.Printf("Stats:\n")
    for _, stat := range val.Stats {
        fmt.Printf("    %v- %v\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, tp := range val.Types {
       fmt.Printf("    -%v\n", tp.Type.Name) 
    }
    return nil
}

// Pokedex function prints all the pokemon in the users's pokedex to the console
func Pokedex(p ...string) error {
    if len(pokedex) <= 0 {
        fmt.Println("No pokemon yet!")
        return nil
    }
    fmt.Println("Your Pokemon:")
    for key := range pokedex {
        fmt.Printf("    -%v\n", key)
    }
    return nil
}
