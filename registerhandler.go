package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kevinjimenez96/gator/internal/database"
)

func RegisterHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	dbuser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		return err
	}

	s.cfg.SetUser(dbuser.Name)

	fmt.Printf("User have been created: %v\n", dbuser)

	return nil
}
