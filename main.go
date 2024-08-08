package main

import ( 
    "fmt"
    "bufio"
    "os"
    "strings"
    "github.com/Greeshmanth1909/pokedexCli/api"
)

type cliCommand struct {
        name string
        description string
        callBack func(str ...string) error
}

func helpCallback(str ...string) error {
    fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
    return nil
}

func exitCallback(str ...string) error {
    fmt.Println("Exiting...")
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
            callBack: api.CommandMap,
        },
        "mapb": {
            name: "mapb",
            description: "displays the names of previous 20 locations, if there are any",
            callBack: api.CommandMapb,
        },
        "explore": {
            name: "explore",
            description: "displays the names of pokemon present in the given area",
            callBack: api.Explore,
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
        if input == "mapb" {
            command["mapb"].callBack()
        }
        // extract multiple inputs, if any
        inputs := strings.Fields(input)
        if inputs[0] == "explore" {
            if len(inputs) != 2 {
                fmt.Println("Please enter a city to explore, explore <city-name>")
                continue
            }
            command["explore"].callBack(inputs[1])
        }
    }

}
