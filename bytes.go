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

// String
func Bytes(str []byte) ParserFunc[[]byte] {
	return func(input []byte) (ParseResult[[]byte], error) {
		if len(input) < len(str) {
			return ParseResult[[]byte]{
				Parsed: nil,
				Remain: input,
			}, &ParseError{
				Cause: "Bytes",
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
				Cause: "Bytes",
				RemainLength: len(input),
			}
		}
	}
}

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
				Cause: "Until",
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
					Cause: errStr,
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
func Space0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isSpace, 0, "space0")
}
func Space1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isSpace, 1, "space1")
}
func MultiSpace0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isMultiSpace, 0, "multiSpace1")
}
func MultiSpace1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isMultiSpace, 1, "multiSpace1")
}
func isAlpha(b byte) bool {
	return 'A' <= b && b <= 'Z' || 'a' <= b && b <= 'z'
}
func isNumeric(b byte) bool {
	return '0' <= b && b <= '9'
}
func Alpha0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlpha, 0, "alpha0")
}
func Alpha1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlpha, 1, "alpha1")
}
func Numeric0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isNumeric, 0, "numeric0")
}
func Numeric1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isNumeric, 1, "numeric1")
}
func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isNumeric(b)
}
func AlphaNumeric0(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlphaNumeric, 0, "alphaNumeric0")
}
func AlphaNumeric1(input []byte) (ParseResult[[]byte], error) {
	return checkByteSequence(input, isAlphaNumeric, 1, "alphaNumeric1")
}
