package main

import (
	"context"
	"fmt"
)

func UsersHandler(s *state, cmd command) error {
	dbusers, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range dbusers {
		if s.cfg.CurrentUsername == u.Name {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}

	}

	return nil
}
