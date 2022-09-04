package godrink

import "fmt"

// ParseError
type ParseError struct {
	Cause string
	RemainLength int
	ParentError *ParseError
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
		if v == '\n' {
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

func GetSpecificLineFromInput(input []byte, targetRow int) []byte {
	res := []byte{}
	row := 0
	for _, v := range input {
		if v == '\n' {
			row += 1
		} else {
			if row == targetRow {
				res = append(res, v)
			}
		}
	}
	return res
}

func FormattedErrorMessage(input []byte, pe *ParseError) string {
	loc := GetErrorLocation(input, pe)
	res := fmt.Sprintf("Parse error at row %d, col %d, caused by %s\n", loc.Row + 1, loc.Col + 1, pe.Cause)
	line := fmt.Sprintf("%d: ", loc.Row + 1)
	res += line
	res += string(GetSpecificLineFromInput(input, loc.Row))
	res += "\n"
	for i := 0; i < len(line) + loc.Col; i++ {
		res += " "
	}
	res += "^"
	if pe.ParentError != nil {
		res += "\nThis error caused by this parent error:\n"
		res += FormattedErrorMessage(input, pe.ParentError)
	}
	return res
}
