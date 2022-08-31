// Package godrink is a simple parser combinator library
// that provides the functions which required for
// implementing a parser.
//
// You can use these functions procedurally but
// we recommends using combine them.
//
// godrink reads the input bytes in order and
// parses them into the desired structure.
// This behavior is inspired by nom,
// the parser combinator library written in Rust.
package godrink

// ParseResult contains a pointer to the parsed object and
// the remaining bytes.
// Parser should return ParseResult and ParseError.
//
// The type parameter T appears in the Parsed field.
// The type of the Parsed field is a pointer of T.
// The Parsed field may be nil (e.g. Optional function
// may return a ParseResult where the Parsed field is nil).
//
// The Remain field is the remainder of the parsed input.
type ParseResult[T any] struct {
	Parsed *T
	Remain []byte
}

// ParserFunc is a function type definition that takes bytes
// as an argument and returns a ParseResult and an error.
// The return type of a function returning a parser should be this.
type ParserFunc[T any] (func([]byte) (ParseResult[T], error))
