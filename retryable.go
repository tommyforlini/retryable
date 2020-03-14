package retry

import (
	"time"
)

// Config represents the retryable options
type Config struct {
	maxAttempts int
	delay       time.Duration
}

// Retryable usable interface function
type retryable interface {
	Try(retryableFunc retryableFunc) error
}

// RetryableFunc represents the function that is attempted to be retryable
type retryableFunc func() error

// Try is the core piece of functionality
func (c *Config) Try(retryableFunc retryableFunc) error {
	var iteration int

	errorLog := make(Error, c.maxAttempts)

	lastErrIndex := iteration

	for iteration < c.maxAttempts {
		err := retryableFunc()

		if err != nil {
			errorLog[lastErrIndex] = parseUnrecoverable(err)
			lastErrIndex++

			// don't retry if flagged as unrecoverable error type
			if !IsRecoverable(err) {
				break
			}

			// don't sleep if we've exhausted all attempts
			if iteration == c.maxAttempts-1 {
				break
			}

			time.Sleep(c.delay)
		} else {
			return nil
		}

		iteration++
	}

	return errorLog
}
