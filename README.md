# pokedexCli
A fun pokemon game which enables you to explore the pokemon world and catch pokemon, all from the command line!

## Installation
The game requires Go 1.22.x or higher.
With the required version of Golang installed, clone the repository with

```
git clone https://github.com/Greeshmanth1909/pokedexCli
```
To compile a binary, run this command in the root of the project ie. the directory that contains `main.go`
```
go build
```
To run the program without building, run
```
go run .
```
Once the program starts running, the terminal screen should have a prompt that says `Pokedex >` and look something like this
<img width="567" alt="Screenshot 2024-08-24 at 11 57 43 AM" src="https://github.com/user-attachments/assets/9c27c390-f52c-4c49-927e-c533ce7da5b5">

Here are the list of commands that can be entered,

|Command|Description|Usage|
|---|---|---|
|help|Displays a help message|`Pokedex >help`
exit|Quits the game|`Pokedex >exit`
map|Displays the names of the next 20 locations in the pokemon world|`Pokedex >map`
mapb|Displays the names of the previous 20 locations in the pokemon world, if there are any|`Pokedex >mapb`
explore|Displays the names of the pokemon present in the provided area|`Pokedex >explore <area-name>`
catch|Attempts to catch a pokemon, the chance of a successful catch depends on the `base_experience` of a pokemon. Higher the `base_xp` the harder it is to catch.| `Pokedex >catch <pokemon-name>`
inspect|Shows important statistics of a caught pokemon|`Pokedex >inspect <pokemon-name>`
pokedex|Displays all the pokemon the user caught|`Pokedex >pokedex`


## Technicals

### API
The program connects to the [pokeApi](https://pokeapi.co/) to fetct the necessary resources

### Caching
The program caches fetched resources in memory for 5 minutes to improve performance
