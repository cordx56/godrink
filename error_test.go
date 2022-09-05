package godrink

import "fmt"

func ExampleGetErrorLocation() {
	input := []byte("abc\n012")
	_, err := Sequence(Alpha1, MultiSpace0, Alpha1)(input)
	loc := GetErrorLocation(input, err.(*ParseError))
	fmt.Printf("Row: %d, Col: %d\n", loc.Row + 1, loc.Col + 1)
	// Output:
	// Row: 2, Col: 1
}

func ExampleFormatErrorMessage() {
	input := []byte("abc\n012")
	_, err := Sequence(Alpha1, MultiSpace0, Alpha1)(input)
	errmsg := FormatErrorMessage(input, err.(*ParseError))
	fmt.Println(errmsg)
	// Output:
	// Parse error at row 2, col 1, caused by Alpha1
	// 2: 012
	//    ^
}
