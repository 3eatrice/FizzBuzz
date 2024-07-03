package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should return a number as string", func(t *testing.T) {

		assert.Equal(t, "1", FizzBuzz(1))
		assert.Equal(t, "2", FizzBuzz(2))
		assert.Equal(t, "4", FizzBuzz(4))
	})

	t.Run("should return Fizz", func(t *testing.T) {

		assert.Equal(t, "Fizz", FizzBuzz(3))
		assert.Equal(t, "Fizz", FizzBuzz(6))
		assert.Equal(t, "Fizz", FizzBuzz(9))
	})

	t.Run("should return Buzz", func(t *testing.T) {

		assert.Equal(t, "Buzz", FizzBuzz(5))
		assert.Equal(t, "Buzz", FizzBuzz(10))
	})

	t.Run("should return FizzBuzz", func(t *testing.T) {

		assert.Equal(t, "FizzBuzz", FizzBuzz(15))
		assert.Equal(t, "FizzBuzz", FizzBuzz(30))
	})

}
