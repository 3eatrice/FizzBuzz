package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		acutal := FizzBuzz()
		expect := "1"

		assert.Equal(t, expect, acutal)
	})
}