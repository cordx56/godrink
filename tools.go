package godrink

import "errors"

// Option
func Option[T any](p ParserFunc[T]) ParserFunc[T] {
	return func(input []byte) (ParseResult[T], error) {
		res, _ := p(input)
		return res, nil
	}
}

// Try
func Try[T any](ps ...ParserFunc[T]) ParserFunc[T] {
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
		}, errors.New("try")
	}
}

// Transform
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

func many[T any](p ParserFunc[T], minCount int, errStr string) ParserFunc[[]T] {
	return func(input []byte) (ParseResult[[]T], error) {
		var result []T
		remain := input
		for i := 0; ; i++ {
			res, err := p(remain)
			if err != nil {
				if i < minCount {
					return ParseResult[[]T]{
						Parsed: nil,
						Remain: res.Remain,
					}, errors.New(errStr)
				} else {
					return ParseResult[[]T]{
						Parsed: &result,
						Remain: remain,
					}, nil
				}
			}
			result = append(result, *res.Parsed)
			remain = res.Remain
		}
	}
}
func Many0[T any](p ParserFunc[T]) ParserFunc[[]T] {
	return many(p, 0, "many0")
}
func Many1[T any](p ParserFunc[T]) ParserFunc[[]T] {
	return many(p, 1, "many1")
}
