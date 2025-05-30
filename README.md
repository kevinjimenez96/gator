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

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database

