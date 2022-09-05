package godrink

import "fmt"

func ExampleInteger() {
	res, _ := Integer([]byte("123abc"))
	fmt.Printf("%d\n%s\n", *res.Parsed, string(res.Remain))
	// Output:
	// 123
	// abc
}

func ExampleFloat() {
	res, _ := Float([]byte("3.14"))
	fmt.Printf("%0.2f\n%s\n", *res.Parsed, string(res.Remain))
	// Output:
	// 3.14
	//
}
