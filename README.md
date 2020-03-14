# retryable

Abstracted way to enable retry mechanisms for any function

## Usage

- Set a `maxAttempts` value which is to be used in the retryable logic
- Set a `delay` value to wait before executing the retryable function

## Sample 1 

### - Anonymous function with no error

> Run 2 attempts with a 5 second delay before retrying.
Final result will be 1 execution with a success.

```golang

	retryable := &Config{
		maxAttempts: 2,
		delay:       time.Second * 5,
	}

	err := retryable.Try(func() error {
        fmt.Println(`
        My cool function will only execute 
        once because no error occurred!`)
		return nil
	})
```

### - Anonymous function with recoverable error

> Run 2 attempts with a 5 second delay before retrying.
Final result will be 2 executions with a failed.

```golang

	retryable := &Config{
		maxAttempts: 2,
		delay:       time.Second * 5,
	}

	err := retryable.Try(func() error {
        fmt.Println(`
        My cool function will execute 2 times
        because an error occurred during each call
        but the final result will be a fail!`)
		return errors.New("forced it to fail")
	})
```

### - Anonymous function with unrecoverable error

> Run 2 attempts with a 5 second delay before retrying.
Final result will be 1 execution with a failed.

```golang

	retryable := &Config{
		maxAttempts: 2,
		delay:       time.Second * 5,
	}

	err := retryable.Try(func() error {
        fmt.Println(`
        My cool function will execute 1 times
        because an UNRECOVERABLE error occurred during 
        the first call but the final result will be a fail!`)
		return Unrecoverable(errors.New("forced it to fail with unrecoverable error"))
	})
```

## Tests

Run test suite

```bash
go test ./...
```
