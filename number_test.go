package godrink

import "fmt"

func ExampleInteger() {
	res, _ := Integer([]byte("123abc"))
	fmt.Printf("%d\n%s\n", *res.Parsed, string(res.Remain))
	// Output:
	// 123
	// abc
}
