package godrink

import "fmt"

// ParseError contains the name of the function that caused the parse error,
// the number of bytes remaining from the point where the error occurred,
// and the parent error, if any.
type ParseError struct {
	Cause string
	RemainLength int
	ParentError *ParseError
}
// Error function returns the name of the function that caused the parse error.
func (pe *ParseError) Error() string {
	return pe.Cause
}

// Location contains row number and col number.
type Location struct {
	Row int
	Col int
}

// GetErrorFunction takes an input and an error as arguments and
// outputs the location of the error as Location.
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

func getSpecificLineFromInput(input []byte, targetRow int) []byte {
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

// FormatErrorMessage function takes an input and an error as
// arguments and returns a formatted error message.
func FormatErrorMessage(input []byte, pe *ParseError) string {
	loc := GetErrorLocation(input, pe)
	res := fmt.Sprintf("Parse error at row %d, col %d, caused by %s\n", loc.Row + 1, loc.Col + 1, pe.Cause)
	line := fmt.Sprintf("%d: ", loc.Row + 1)
	res += line
	res += string(getSpecificLineFromInput(input, loc.Row))
	res += "\n"
	for i := 0; i < len(line) + loc.Col; i++ {
		res += " "
	}
	res += "^"
	if pe.ParentError != nil {
		res += "\nThis error caused by this parent error:\n"
		res += FormatErrorMessage(input, pe.ParentError)
	}
	return res
}
