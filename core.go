package godrink

// ParseResult contains parsed object and remain bytes
type ParseResult[T any] struct {
	Parsed *T
	Remain []byte
}

// ParserFunc is function definition that takes bytes as argument and return ParseResult and error
type ParserFunc[T any] (func([]byte) (ParseResult[T], error))
