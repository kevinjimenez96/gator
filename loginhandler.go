package main

import (
	"context"
	"fmt"
)

func LoginHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	dbuser, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("user not found: %s", cmd.arguments[0])
	}

	err = s.cfg.SetUser(dbuser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("Username have been set to " + cmd.arguments[0])

	return nil
}
