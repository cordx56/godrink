package godrink

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
