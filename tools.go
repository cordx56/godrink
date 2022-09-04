package godrink

// Optional function succeeds either the function
// given as the argument succeeds or not.
// Optional takes ParserFunc[T] as an argument and
// returns a ParserFunc[T].
func Optional[T any](p ParserFunc[T]) ParserFunc[T] {
	return func(input []byte) (ParseResult[T], error) {
		res, _ := p(input)
		return res, nil
	}
}
// Opt function is alias of the Optional function
func Opt[T any](p ParserFunc[T]) ParserFunc[T] {
	return Optional(p)
}

// Any function tries functions given as arguments in order.
func Any[T any](ps ...ParserFunc[T]) ParserFunc[T] {
	return func(input []byte) (ParseResult[T], error) {
		for _, p := range ps {
			res, err := p(input)
			if err == nil {
				return res, nil
			}
		}
		return ParseResult[T]{
				Parsed: nil,
				Remain: input,
			}, &ParseError{
				Cause:        "Any",
				RemainLength: len(input),
			}
	}
}

// Transform function transforms the parsed result to another type
// using the transformer function given as an argument.
func Transform[T any, U any](p ParserFunc[T], transformer func(T) U) ParserFunc[U] {
	return func(input []byte) (ParseResult[U], error) {
		res, err := p(input)
		if err != nil {
			return ParseResult[U]{
				Parsed: nil,
				Remain: res.Remain,
			}, err
		}
		mapped := transformer(*res.Parsed)
		return ParseResult[U]{
			Parsed: &mapped,
			Remain: res.Remain,
		}, nil
	}
}
// Tf function is alias of the Transform function
func Tf[T any, U any](p ParserFunc[T], transformer func(T) U) ParserFunc[U] {
	return Transform(p, transformer)
}

func many[T any](p ParserFunc[T], minCount int, errStr string) ParserFunc[[]T] {
	return func(input []byte) (ParseResult[[]T], error) {
		var result []T
		remain := input
		for i := 0; ; i++ {
			remainLen := len(remain)
			res, err := p(remain)
			if err != nil {
				if i < minCount {
					return ParseResult[[]T]{
							Parsed: nil,
							Remain: input,
						}, &ParseError{
							Cause:        errStr,
							RemainLength: len(res.Remain),
							ParentError:  err.(*ParseError),
						}
				} else {
					return ParseResult[[]T]{
						Parsed: &result,
						Remain: remain,
					}, nil
				}
			} else {
				// Infinite loop
				if remainLen == len(res.Remain) {
					return ParseResult[[]T]{
							Parsed: nil,
							Remain: input,
						}, &ParseError{
							Cause:        errStr,
							RemainLength: len(input),
						}
				}
			}
			result = append(result, *res.Parsed)
			remain = res.Remain
		}
	}
}
// Many0 function repeatedly attempts to parse with
// the function passed as an argument.
func Many0[T any](p ParserFunc[T]) ParserFunc[[]T] {
	return many(p, 0, "Many0")
}
// Many1 function repeatedly attempts to parse with
// the function passed as an argument.
//
// Many1 function require at least 1 matching.
func Many1[T any](p ParserFunc[T]) ParserFunc[[]T] {
	return many(p, 1, "Many1")
}

// NoRemain function forces the function passed as
// an argument to consume all input.
func NoRemain[T any](p ParserFunc[T]) ParserFunc[T] {
	return func(input []byte) (ParseResult[T], error) {
		res, err := p(input)
		if 0 < len(res.Remain) {
			if err != nil {
				return ParseResult[T]{
						Parsed: res.Parsed,
						Remain: input,
					}, &ParseError{
						Cause:        "NoRemain",
						RemainLength: len(res.Remain),
						ParentError:  err.(*ParseError),
					}
			} else {
				return ParseResult[T]{
						Parsed: res.Parsed,
						Remain: input,
					}, &ParseError{
						Cause:        "NoRemain",
						RemainLength: len(res.Remain),
						ParentError:  nil,
					}
			}
		}
		return res, err
	}
}

// Not function returns an error if the parser passed as
// an argument succeeds, otherwise it returns success.
func Not[T any](p ParserFunc[T]) ParserFunc[[]byte] {
	return func(input []byte) (ParseResult[[]byte], error) {
		_, err := p(input)
		if err != nil {
			return ParseResult[[]byte]{
				Parsed: &input,
				Remain: []byte{},
			}, nil
		} else {
			return ParseResult[[]byte]{
					Parsed: nil,
					Remain: input,
				}, &ParseError{
					Cause:        "Not",
					RemainLength: len(input),
				}
		}
	}
}
