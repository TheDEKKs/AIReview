package handlers

import "thedekk/AIReview/internal/api"

func Request() error {
	if err := api.Test(); err != nil {
		return err
	}
	return nil
}
