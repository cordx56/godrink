package godrink

type Location struct {
	Row int
	Col int
}

func GetErrorLocation[T any](input []byte, res ParseResult[T]) Location {
	row := 0
	col := 0
	for _, v := range input[:len(input) - len(res.Remain)] {
		if v == 10 {
			row += 1
			col = 0
		} else {
			col += 1
		}
	}
	return Location{
		Row: row,
		Col: col,
	}
}
