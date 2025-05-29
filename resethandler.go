package main

import (
	"context"
	"fmt"
)

func ResetHandler(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Users have been delted. \n")

	return nil
}
