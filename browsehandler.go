package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kevinjimenez96/gator/internal/database"
)

func BrowseHandler(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.arguments[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		Limit:    int32(limit),
		Username: user.Name,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title.String)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
