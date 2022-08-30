package godrink

import (
	"bytes"
	"errors"
	"fmt"
)

// String
func String(tag []byte) ParserFunc[[]byte] {
	return func(input []byte) (ParseResult[[]byte], error) {
		if len(input) < len(tag) {
			return ParseResult[[]byte]{
				Parsed: nil,
				Remain: input,
			}, errors.New(fmt.Sprintf("tag(%s)", tag))
		}
		if bytes.Equal(tag, input[:len(tag)]) {
			return ParseResult[[]byte]{
				Parsed: &tag,
				Remain: input[len(tag):],
			}, nil
		} else {
			return ParseResult[[]byte]{
				Parsed: nil,
				Remain: input,
			}, errors.New(fmt.Sprintf("tag(%s)", tag))
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
				}, errors.New(errStr)
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
