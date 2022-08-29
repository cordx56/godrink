package godrink

// Sequence returns a function that takes input as argument and returns parse result.
// The function returns by Sequesnce function tries to parse input in order of Sequence function's arguments.
// If all parsers passed as arguments succeeds to parse input, then return *ParseResult[[]T] and nil.
// If one of the parsers failed to parse input, then return nil and error.
func Sequence[T any](ps ...ParserFunc[T]) ParserFunc[[]T] {
	return func(input []byte) (ParseResult[[]T], error) {
		result := []T{}
		remain := input
		for _, p := range ps {
			res, err := p(remain)
			if err != nil {
				return ParseResult[[]T]{
					Parsed: nil,
					Remain: remain,
				}, err
			}
			result = append(result, res.Parsed)
			remain = res.Remain
		}
		return ParseResult[[]T]{
			Parsed: result,
			Remain: remain,
		}, nil
	}
}

// Pair is a struct that contains pair of datum
type Pair[T any, U any] struct {
	Prev T
	Next U
}

// Next
func Next[T any, U any](p0 ParserFunc[T], p1 ParserFunc[U]) ParserFunc[*Pair[T, U]] {
	return func(input []byte) (ParseResult[*Pair[T, U]], error) {
		r0, err := p0(input)
		if err != nil {
			return ParseResult[*Pair[T, U]]{
				Parsed: nil,
				Remain: input,
			}, err
		}

		r1, err := p1(r0.Remain)
		if err != nil {
			return ParseResult[*Pair[T, U]]{
				Parsed: nil,
				Remain: input,
			}, err
		}

		return ParseResult[*Pair[T, U]]{
			Parsed: &Pair[T, U]{
				Prev: r0.Parsed,
				Next: r1.Parsed,
			},
			Remain: r1.Remain,
		}, nil
	}
}
