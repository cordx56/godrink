# godrink - Simple parser combinator

godrink is a simple parser combinator written in Golang, inspired by [nom](https://github.com/Geal/nom) written in Rust.

## Feature
godrink is simple, lightweight library.
godrink does not depend on non-standard libraries.

## Example
### Example of function use
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
				godrink.Try(
					godrink.String([]byte("+")),
					godrink.String([]byte("-")),
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
		fmt.Println(godrink.FormattedErrorMessage(inputBytes, err.(*godrink.ParseError)))
	} else {
		// Display calculation result
		fmt.Println(*res.Parsed)
	}
}
```
