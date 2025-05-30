package main

import (
	"context"
	"fmt"

	"github.com/kevinjimenez96/gator/internal/database"
)

func UnfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Username: %s\n unfollowed Feed: %s.\n", user.Name, feed.Name)

	return nil
}
