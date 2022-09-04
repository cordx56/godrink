package godrink

import "testing"

func TestUntil(t *testing.T) {
	parsedBytes1 := []byte("012")
	testParser(t, parserTestDefinition[[]byte]{
		name: "Take until",
		parser: Until([]byte("a")),
		input: []byte("012abc"),
		res: ParseResult[[]byte]{
			Parsed: &parsedBytes1,
			Remain: []byte("abc"),
		},
		err: nil,
	})
}
