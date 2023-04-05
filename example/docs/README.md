# Command examples for redis.io

## How to add examples

- Create a Go test file with a meaningful name in the current folder. 
- Create a single method prefixed with `Example` and write your test in it.
- We're using the [Testable Examples](https://go.dev/blog/examples) feature of Go to test the desired output has been written to stdout.
- To run the tests start a Redis server locally on port 6379 and run `go test`

### Special markup

#### HIDE_START and HIDE_END
Should be used to hide imports, connection creation and other bootstrapping code that is not important
for understanding a command example.

Example:

```go
// PAGES: Set, Get
// HIDE_START
package example_docs_test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ExampleSetGet() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password docs
		DB:       0,  // use default DB
	})

	// HIDE_END
	err := rdb.Set(ctx, "bike:1", "Process 134", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("OK")
}

```

#### REMOVE_START and REMOVE_END
Should be used to **remove** code from the resulting code snippet published on redis.io.
This markup can be useful to remove assertions, and any eventual other testing blocks.

## How to test the examples

- Start a Redis server locally on port 6379
- CD into the `example/docs` directory
- Run `go test` to test all examples in the current folder.
- Run `go test filename.go` to test a single file

## Missing

- We need to set up a way to automatically bring up a Redis server on port 6379 before the test suite, flush the db between tests and stop the server after the test suite.
- We should be able to test with more than one Redis version