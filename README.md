# godrink - simple parser combinators library

godrink is a simple parser combinators library written in Golang, inspired by [nom](https://github.com/Geal/nom) written in Rust.

## Feature
godrink is a simple, lightweight library.
godrink does not depend on non-standard libraries.

godrink reads input bytes and applies parsers in sequence and parses them into the desired structures.

## Documentation
See [pkg.go.dev/github.com/cordx56/godrink](https://pkg.go.dev/github.com/cordx56/godrink).

## Example
### Example projects
See these repositories.
- [cordx56/gomdrink](https://github.com/cordx56/gomdrink), a markdown parser written using godrink

### Example code
```golang
package main

import (
	"bytes"
	"fmt"

	"github.com/cordx56/godrink"
)

func main() {
	// input expression
	input := "1+2"

	// Transform parsed result to integer
	parser := godrink.Transform(
		// Parse integer and operator in order
		godrink.Next(
			godrink.Integer,
			godrink.Next(
				godrink.Any(
					godrink.Bytes([]byte("+")),
					godrink.Bytes([]byte("-")),
				),
				godrink.Integer,
			),
		),
		// Transformer to transform from parsed struct to integer
		func(parsed godrink.Pair[int, godrink.Pair[[]byte, int]]) int {
			left := parsed.Prev
			right := parsed.Next.Next
			if bytes.Equal(parsed.Next.Prev, []byte("+")) {
				return left + right
			} else {
				return left - right
			}
		},
	)

	inputBytes := []byte(input)
	// perform perse
	res, err := parser(inputBytes)
	if err != nil {
		// Display Error message
		fmt.Println(godrink.FormatErrorMessage(inputBytes, err.(*godrink.ParseError)))
	} else {
		// Display calculation result
		fmt.Println(*res.Parsed)
	}
}
```
