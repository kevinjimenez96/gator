package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kevinjimenez96/gator/internal/config"
	"github.com/kevinjimenez96/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

var appState = state{}
var cmds = commands{
	commands: make(map[string]func(*state, command) error),
}

func repl() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal("db error!")
	}
	dbQueries := database.New(db)

	appState.cfg = &cfg
	appState.db = dbQueries

	cmds.register("login", LoginHandler)
	cmds.register("register", RegisterHandler)
	cmds.register("reset", ResetHandler)
	cmds.register("users", UsersHandler)
	cmds.register("agg", AggHandler)
	cmds.register("addfeed", middlewareLoggedIn(AddFeedHandler))
	cmds.register("feeds", FeedsHandler)
	cmds.register("follow", middlewareLoggedIn(FollowHandler))
	cmds.register("unfollow", middlewareLoggedIn(UnfollowHandler))
	cmds.register("following", middlewareLoggedIn(FollowingHandler))
	cmds.register("browse", middlewareLoggedIn(BrowseHandler))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	fmt.Println(os.Args)

	var cmd = command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
