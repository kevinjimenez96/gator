# gator
## Prerequisites

To use **gator**, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or newer)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the gator CLI with:

```sh
go install https://github.com/kevinjimenez96/gator@latest
```

## Configuration

Create a configuration file named `.gatorconfig.json` in your home directory. Example:

```json
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"username"}
```

Update the values as needed for your environment.

## Usage

Run the program with:

```sh
gator
```

Some useful commands:

- `gator register <username>` – Add new username.
- `gator login <username>` – Logs-in with username.
- `gator addfeed <name> <url>` – Add a new rss feed
