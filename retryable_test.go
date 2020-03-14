package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {

	config := &Config{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := config.Try(func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestSuccessExternalizedFunction(t *testing.T) {

	config := &Config{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := config.Try(func() error {
		return MyCoolFunction()
	})

	assert.NoError(t, err)
}

func MyCoolFunction() error {
	return nil
}

func TestMaxAttemptsExhausted(t *testing.T) {

	config := &Config{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := config.Try(func() error {
		return errors.New("It failed")
	})

	assert.Error(t, err)
	expectedErrorFormat := `Retryable failed:
#1: It failed
#2: It failed`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error messages")

}

func TestUnrecoverableError(t *testing.T) {

	config := &Config{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := config.Try(func() error {
		return Unrecoverable(errors.New("It failed with unrecoverable error"))
	})

	assert.Error(t, err)
	expectedErrorFormat := `Retryable failed:
#1: It failed with unrecoverable error`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error messages")

}
