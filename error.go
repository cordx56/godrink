package godrink

// ParseError
type ParseError struct {
	Cause string
	RemainLength int
}
func (pe *ParseError) Error() string {
	return pe.Cause
}

type Location struct {
	Row int
	Col int
}

func GetErrorLocation(input []byte, pe *ParseError) Location {
	row := 0
	col := 0
	for _, v := range input[:len(input) - pe.RemainLength] {
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
