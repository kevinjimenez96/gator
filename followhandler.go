package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kevinjimenez96/gator/internal/database"
)

func FollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Username: %s\nFeed: %s.\n", user.Name, feed.Name)

	return nil
}
