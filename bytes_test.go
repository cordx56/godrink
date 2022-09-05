package godrink

import "fmt"

func ExampleBytes() {
	res, _ := Bytes([]byte("abc"))([]byte("abc123"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// abc
	// 123
}

func ExampleUntil() {
	res, _ := Until([]byte("12"))([]byte("abc123"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// abc
	// 123
}

func ExampleSpace0() {
	res, _ := Space0([]byte("abc"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// abc
}

func ExampleSpace1() {
	res, _ := Space1([]byte(" \tabc"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// abc
}

func ExampleMultiSpace0() {
	res, _ := MultiSpace0([]byte("abc"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// abc
}

func ExampleMultiSpace1() {
	res, _ := MultiSpace1([]byte(" \n abc"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	//
	// abc
}

func ExampleAlpha0() {
	res, _ := Alpha0([]byte("012"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// 012
}

func ExampleAlpha1() {
	res, _ := Alpha1([]byte("abc012"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// abc
	// 012
}

func ExampleNumeric0() {
	res, _ := Numeric0([]byte("abc012"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// abc012
}

func ExampleNumeric1() {
	res, _ := Numeric1([]byte("012abc"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// 012
	// abc
}

func ExampleAlphaNumeric0() {
	res, _ := AlphaNumeric0([]byte("テスト"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	//
	// テスト
}

func ExampleAlphaNumeric1() {
	res, _ := AlphaNumeric1([]byte("abc012"))
	fmt.Printf("%s\n%s\n", string(*res.Parsed), string(res.Remain))
	// Output:
	// abc012
	//
}
