package godrink

import (
	"fmt"
	"testing"
)

func ExampleOptional() {
	res, _ := Optional(Bytes([]byte("456")))([]byte("abc123"))
	fmt.Printf("%v\n%s\n", res.Parsed, string(res.Remain))
	// Output:
	// <nil>
	// abc123
}

func ExampleAny() {
	res, _ := Any(Bytes([]byte("123")), Bytes([]byte("abc")))([]byte("abc123"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// abc
	// 123
}

func ExampleTransform() {
	res, _ := Transform(
		Integer,
		func(parsed int) int {
			return parsed + 1
		},
	)([]byte("123"))
	fmt.Printf("%d\n%s\n", *res.Parsed, string(res.Remain))
	// Output:
	// 124
	//
}

func ExampleMany0() {
	res, _ := Many0(Bytes([]byte("123")))([]byte("abc123"))
	fmt.Printf("%d\n%s\n", len(*res.Parsed), string(res.Remain))
	// Output:
	// 0
	// abc123
}

func ExampleMany1() {
	res, _ := Many1(Bytes([]byte("abc")))([]byte("abcabc123"))
	fmt.Printf("%d\n%s\n", len(*res.Parsed), string(res.Remain))
	// Output:
	// 2
	// 123
}

func ExampleNoRemain() {
	res, err := NoRemain(Bytes([]byte("abc")))([]byte("abc123"))
	fmt.Printf("%s\n%s\n", string(res.Remain), err.(*ParseError).Cause)
	// Output:
	// abc123
	// NoRemain
}

func ExampleNot() {
	res, err := Not(Bytes([]byte("123")))([]byte("abc123"))
	fmt.Printf("%s\n%v\n", string(*res.Parsed), err)
	// Output:
	// abc123
	// <nil>
}


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
