package main

import ( 
    "fmt"
    "bufio"
    "os"
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
    }

}
