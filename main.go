package main

import ( 
    "fmt"
    "bufio"
    "os"
    "net/http"
    "io/ioutil"
)

type cliCommand struct {
        name string
        description string
        callBack func() error
}

func helpCallback() error {
    fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
    return nil
}

func exitCallback() error {
    fmt.Println("Exiting...")
    return nil
}

func commandMap() error {
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
    fmt.Println(string(body))
    return nil

}

func main() {
    // get user input from console
    
    command := map[string] cliCommand{
        "help": {
            name: "help",
            description: "Displays a help message",
            callBack: helpCallback,
        },
        "exit": {
            name: "exit",
            description: "quits the REPL",
            callBack: exitCallback,
        },
        "map": {
            name: "map",
            description: "displays the names of 20 location areas in the Pokemon world",
            callBack: commandMap,
        },

    }

    for true {
        fmt.Print("Pokedex >")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()
        if input == "help" {
            command["help"].callBack()
        }
        if input == "exit" {
            command["exit"].callBack()
            return
        }
        if input == "map" {
            command["map"].callBack()
        }
    }

}
