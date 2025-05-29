package main

import (
	"context"
	"fmt"

	"github.com/kevinjimenez96/gator/internal/rss"
)

func AggHandler(s *state, cmd command) error {
	rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)

	return nil
}
