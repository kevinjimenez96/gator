package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kevinjimenez96/gator/internal/database"
	"github.com/kevinjimenez96/gator/internal/rss"
)

func AggHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <time>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: use 1s, 1h")
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.arguments[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting next feed: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), next.Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true, // Set to true to indicate the time is valid
		},
		FeedID: next.ID,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error marking feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		publishedAt, _ := time.Parse(time.RFC1123Z, item.PubDate)

		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  true,
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      next.ID,
		})
	}

	fmt.Println("")

}
