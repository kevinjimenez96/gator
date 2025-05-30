package main

import (
	"context"
	"fmt"

	"github.com/kevinjimenez96/gator/internal/database"
)

func FeedsHandler(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		printFeed(feed)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeed(feed database.GetAllFeedsRow) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", feed.Username)
}
