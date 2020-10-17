package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"function/src"
)

func TestSuccess(t *testing.T) {
	eventSuccess := src.Event{
		Name:   "something",
		Region: "eu-west-1",
	}

	err := application(context.Background(), eventSuccess)

	assert.Nil(t, err)
}

func TestFailure(t *testing.T) {
	err := application(context.Background(), src.Event{})

	assert.NotNil(t, err)
}
