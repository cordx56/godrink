package godrink

import (
	"reflect"
	"testing"
)

type parserTestDefinition[T any] struct{
	name string
	parser ParserFunc[T]
	input []byte
	res ParseResult[T]
	err error
}

func testParser[T any](t *testing.T, test parserTestDefinition[T]) {
	t.Run(test.name, func(t *testing.T) {
		res, err := test.parser(test.input)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("Returned result %v not matched!\nExpected: %v", res, test.res)
		}
		if err != nil {
			if !reflect.DeepEqual(err.(*ParseError), test.err.(*ParseError)) {
				t.Errorf("Returned error %s not matched!\nExpected: %s", err.Error(), test.err.Error())
			}
		} else {
			if err != test.err {
				t.Errorf("Returned error nil not matched!\nExpected: %s", test.err.Error())
			}
		}
	})
}
