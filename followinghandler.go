package main

import (
	"context"
	"fmt"

	"github.com/kevinjimenez96/gator/internal/database"
)

func FollowingHandler(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Println("Your are following:")
	for _, feed := range feeds {
		fmt.Printf("  - %s\n", feed.FeedName)
	}

	return nil
}
