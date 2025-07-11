# Pokédex CLI

A command-line Pokédex application written in Go.  
Explore, catch, and inspect Pokémon using the PokéAPI, with caching for faster lookups.

---

## Features

- **Explore** the Pokémon world and retrieve Pokémon details.  
- **Catch** Pokémon and add them to your personal Pokédex.  
- **Inspect** caught Pokémon.  
- **Map / MapB** to discover available locations.  
- **Pokedex** to view your collection.  
- **Caching** of API calls to speed up repeated lookups.

---

## Commands

Below is an overview of the available commands:

| Command   | Description                                         |
|-----------|-----------------------------------------------------|
| `help`    | Show help information and list available commands.  |
| `explore` | Explore a location or list available Pokémon.       |
| `catch`   | Catch a Pokémon by name.                            |
| `inspect` | View detailed information about a caught Pokémon.   |
| `map`     | Show a list of all available Pokémon locations.     |
| `mapb`    | Show more detailed location information.            |
| `pokedex` | List all Pokémon you have caught so far.            |
| `exit`    | Exit the CLI.                                       |

---

## Caching

This CLI includes a caching layer for API responses.  
When you query or catch a Pokémon, its data is stored in memory.  
Subsequent calls to the same resource will use the cached response, reducing API requests and improving speed.

The cache expires entries after a configurable duration (e.g., 30 seconds).

---

## Installation

### Clone the repository

```bash
git clone https://github.com/your-username/pokedex-cli.git
cd pokedex-cli
```

### Build the application

```bash
go build -o pokedex
```

### Run the CLI

```bash
./pokedex
```

---

## Example Usage

```text
Pokedex > help
Available commands:
 - explore
 - catch <pokemon>
 - inspect <pokemon>
 - map
 - mapb
 - pokedex
 - help
 - exit

Pokedex > catch pikachu
Throwing a Pokéball at pikachu...
pikachu was caught!

Pokedex > pokedex
Your Pokédex:
pikachu

Pokedex > inspect pikachu
Name: pikachu
Base Experience: 112
Height: 4
Weight: 60
Stats:
 - speed: 90
 - special-defense: 50
 - special-attack: 50
 - defense: 40
 - attack: 55
 - hp: 35
```

---

## Dependencies

- [Go](https://golang.org/) (version 1.18 or newer recommended)  
- [PokéAPI](https://pokeapi.co/) — free Pokémon API

---

## License

This project is licensed under the MIT License.
