package src

import (
	"errors"
	"log"
)

func HandleEvent(payload Event) error {
	log.Printf("%+v", payload)

	if payload.Name == "" {
		return errors.New("empty name")
	}

	return nil
}
