// Package godrink is a simple parser combinators library
// that provides the functions which required for
// implementing a parser.
//
// godrink reads input bytes and applies parsers
// in sequence and parses them into the desired structures.
// This behavior is inspired by nom,
// the parser combinators library written in Rust.
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
