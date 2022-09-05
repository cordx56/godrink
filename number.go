package godrink

import "math"

func bytesToInt(input []byte) int {
	res := 0
	for _, v := range input {
		res *= 10
		res += int(v - '0')
	}
	return res
}

// Integer parses an input integer and returns it as a value of type int.
func Integer(input []byte) (ParseResult[int], error) {
	return Transform(
		Numeric1,
		bytesToInt,
	)(input)
}

// Float parses a float number and returns it as a value of type float64.
func Float(input []byte) (ParseResult[float64], error) {
	return Transform(
		Next(
			Integer,
			Next(
				Bytes([]byte(".")),
				Numeric1,
			),
		),
		func(parsed Pair[int, Pair[[]byte, []byte]]) float64 {
			res := float64(parsed.Prev)
			decimalBytes := parsed.Next.Next
			decimal := float64(bytesToInt(decimalBytes)) / math.Pow10(len(decimalBytes))
			return res + decimal
		},
	)(input)
}
