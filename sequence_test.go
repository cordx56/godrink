package godrink

import "fmt"

func ExampleSequence() {
	res, _ := Sequence(Alpha1, Numeric1)([]byte("abc123"))
	fmt.Printf("%s\n%s\n%s\n", string((*res.Parsed)[0]), string((*res.Parsed)[1]), string(res.Remain))
	// Output:
	// abc
	// 123
	//
}

func ExampleNext() {
	res, _ := Next(Alpha1, Integer)([]byte("abc123"))
	fmt.Printf("%s\n%d\n%s\n", string(*res.Parsed.Prev), *res.Parsed.Next, string(res.Remain))
	// Output:
	// abc
	// 123
	//
}
