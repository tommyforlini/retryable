package retryable

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {

	retryable := &Client{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := retryable.Try(func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestSuccessWithDefault(t *testing.T) {

	err := NewClient().Try(func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestSuccessExternalizedFunction(t *testing.T) {

	retryable := &Client{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := retryable.Try(func() error {
		return MyCoolFunction()
	})

	assert.NoError(t, err)
}

func MyCoolFunction() error {
	return nil
}

func TestMaxAttemptsExhausted(t *testing.T) {

	retryable := &Client{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := retryable.Try(func() error {
		return errors.New("It failed")
	})

	assert.Error(t, err)
	expectedErrorFormat := `Retryable failed:
#1: It failed
#2: It failed`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error messages")

}

func TestMaxAttemptsExhaustedWithDefaultsDelayOverride(t *testing.T) {

	retryable := NewClient()
	retryable.delay = time.Second * 1

	err := retryable.Try(func() error {
		return errors.New("It failed")
	})

	assert.Error(t, err)
	expectedErrorFormat := `Retryable failed:
#1: It failed
#2: It failed
#3: It failed
#4: It failed
#5: It failed`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error messages")

}

func TestUnrecoverableError(t *testing.T) {

	retryable := &Client{
		maxAttempts: 2,
		delay:       time.Second * 1,
	}

	err := retryable.Try(func() error {
		return Unrecoverable(errors.New("It failed with unrecoverable error"))
	})

	assert.Error(t, err)
	expectedErrorFormat := `Retryable failed:
#1: It failed with unrecoverable error`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error messages")

}
