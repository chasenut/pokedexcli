# Pokedex cli

Pokedex is a command-line REPL tool to fetch data from [PokéAPI](https://pokeapi.co/) 
using **go**'s `http` standard library.
Featuring efficient auto-cleanup **cache**, where API responses are stored temporarly 
for later reuse.

## Run and build

To run, use command: 
```
go run .
```

To build the project and make generate executables, use command:
```
go build .
```

If you want to test packages and mechanics, use:
```
go test .
```

## About

Pokedex is a **REPL** type CLI tool that allows you to explore the gigantic world of Pokemon 
using **PokeAPI**. All data is fetched using go's `http` and stored temporarly in **cache**.

Discover new locations on the `map` (`mapb`), `explore` regions, 
`catch` Pokemons and then `inspect` them, and build your own `pokedex`.

## Credit

Made as a project in [boot.dev](https://www.boot.dev) course.
