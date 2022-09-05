package godrink

import "bytes"

func checkBytesEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Bytes function tries to parse a sequence of bytes passed as an argument.
func Bytes(str []byte) ParserFunc[[]byte] {
	return func(input []byte) (ParseResult[[]byte], error) {
		if len(input) < len(str) {
			return ParseResult[[]byte]{
					Parsed: nil,
					Remain: input,
				}, &ParseError{
					Cause:        "Bytes",
					RemainLength: len(input),
				}
		}
		if checkBytesEqual(str, input[:len(str)]) {
			return ParseResult[[]byte]{
				Parsed: &str,
				Remain: input[len(str):],
			}, nil
		} else {
			return ParseResult[[]byte]{
					Parsed: nil,
					Remain: input,
				}, &ParseError{
					Cause:        "Bytes",
					RemainLength: len(input),
				}
		}
	}
}

// Until function returns the input until the first byte sequence
// occurrence specified in the argument.
func Until(strs ...[]byte) ParserFunc[[]byte] {
	return func(input []byte) (ParseResult[[]byte], error) {
		pos := -1
		for _, str := range strs {
			v := bytes.Index(input, str)
			if v != -1 && (pos == -1 || v < pos) {
				pos = v
			}
		}
		if -1 < pos {
			parsed := input[:pos]
			return ParseResult[[]byte]{
				Parsed: &parsed,
				Remain: input[pos:],
			}, nil
		} else {
			return ParseResult[[]byte]{
					Parsed: nil,
					Remain: input,
				}, &ParseError{
					Cause:        "Until",
					RemainLength: len(input),
				}
		}
	}
}

func checkByteSequence(input []byte, checkFunc func(byte) bool, minLength int, errStr string) (ParseResult[[]byte], error) {
	ret := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		if checkFunc(input[i]) {
			ret = append(ret, input[i])
		} else {
			if i < minLength {
				return ParseResult[[]byte]{
						Parsed: nil,
						Remain: input,
					}, &ParseError{
						Cause:        errStr,
						RemainLength: len(input) - len(ret),
					}
			} else {
				return ParseResult[[]byte]{
					Parsed: &ret,
					Remain: input[i:],
				}, nil
			}
		}
	}
	return ParseResult[[]byte]{
		Parsed: &ret,
		Remain: []byte{},
	}, nil
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t'
}
func isMultiSpace(b byte) bool {
	return isSpace(b) || b == '\n' || b == '\r'
}

// Space0 matches zero or more spaces (" " or "\t").
func Space0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isSpace, 0, "Space0")
}

// Space0 matches one or more spaces (" " or "\t").
func Space1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isSpace, 1, "Space1")
}

// MultiSpace0 matches zero or more spaces and line breaks
// (" " or "\t" or "\n" or "\r").
func MultiSpace0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isMultiSpace, 0, "MultiSpace1")
}

// MultiSpace1 matches one or more spaces and line breaks
// (" " or "\t" or "\n" or "\r").
func MultiSpace1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isMultiSpace, 1, "MultiSpace1")
}
func isAlpha(b byte) bool {
	return 'A' <= b && b <= 'Z' || 'a' <= b && b <= 'z'
}
func isNumeric(b byte) bool {
	return '0' <= b && b <= '9'
}

// Alpha0 function matches alphabets of zero or more characters.
func Alpha0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlpha, 0, "Alpha0")
}

// Alpha1 function matches alphabets of one or more characters.
func Alpha1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlpha, 1, "Alpha1")
}

// Numeric0 function matches numeric sequences of zero or more characters.
func Numeric0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isNumeric, 0, "Numeric0")
}

// Numeric1 function matches numeric sequences of one or more characters.
func Numeric1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isNumeric, 1, "Numeric1")
}
func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isNumeric(b)
}

// AlphaNumeric0 function matches a sequence of zero or more
// alphabetic or numeric characters.
func AlphaNumeric0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlphaNumeric, 0, "AlphaNumeric0")
}

// AlphaNumeric1 function matches a sequence of one or more
// alphabetic or numeric characters.
func AlphaNumeric1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlphaNumeric, 1, "AlphaNumeric1")
}
