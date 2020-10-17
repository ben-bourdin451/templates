package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainLogicError(t *testing.T) {
	err := HandleEvent(Event{})

	assert.NotNil(t, err)
}

func TestFailure(t *testing.T) {
	eventSuccess := Event{
		Name:   "something",
		Region: "eu-west-1",
	}

	err := HandleEvent(eventSuccess)

	assert.Nil(t, err)
}
