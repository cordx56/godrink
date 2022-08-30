package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cordx56/godrink"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		s.Scan()
		input := []byte(s.Text())
		res, err := godrink.NoRemain(expr)(input)
		if err != nil {
			pe := err.(*godrink.ParseError)
			errorMessage := godrink.FormattedErrorMessage(input, pe)
			fmt.Println(errorMessage)
		} else {
			fmt.Printf("%f\n", *res.Parsed)
		}
	}
}


// transform numbers into float64
func number(input []byte) (godrink.ParseResult[float64], error) {
	res, err := godrink.Integer(input)
	if err != nil {
		return godrink.ParseResult[float64]{
			Parsed: nil,
			Remain: res.Remain,
		}, err
	} else {
		float := float64(*res.Parsed)
		return godrink.ParseResult[float64]{
			Parsed: &float,
			Remain: res.Remain,
		}, nil
	}
}
// parse number and expression
func factor(input []byte) (godrink.ParseResult[float64], error) {
	return godrink.Try(
		number,
		godrink.Transform(
			godrink.Next(
				godrink.String([]byte("(")),
				godrink.Next(
					godrink.MultiSpace0,
					godrink.Next(
						expr,
						godrink.Next(
							godrink.MultiSpace0,
							godrink.String([]byte(")")),
						),
					),
				),
			),
			func(parsed godrink.Pair[[]byte, godrink.Pair[[]byte, godrink.Pair[float64, godrink.Pair[[]byte, []byte]]]]) float64 {
				return parsed.Next.Next.Prev
			},
		),
	)(input)
}
// parse multiplication and division
func term(input []byte) (godrink.ParseResult[float64], error) {
	return godrink.Transform(
		godrink.Next(
			factor,
			godrink.Many0(
				godrink.Next(
					godrink.MultiSpace0,
					godrink.Next(
						godrink.Try(
							godrink.String([]byte("*")),
							godrink.String([]byte("/")),
						),
						godrink.Next(
							godrink.MultiSpace0,
							factor,
						),
					),
				),
			),
		),
		func(parsed godrink.Pair[float64, []godrink.Pair[[]byte, godrink.Pair[[]byte, godrink.Pair[[]byte, float64]]]]) float64 {
			val := parsed.Prev
			for _, v := range parsed.Next {
				if v.Next.Prev[0] == '*' {
					val *= v.Next.Next.Next
				} else if v.Next.Prev[0] == '/' {
					val /= v.Next.Next.Next
				}
			}
			return val
		},
	)(input)
}
// parse expression
func expr(input []byte) (godrink.ParseResult[float64], error) {
	return godrink.Transform(
		godrink.Next(
			term,
			godrink.Many0(
				godrink.Next(
					godrink.MultiSpace0,
					godrink.Next(
						godrink.Try(
							godrink.String([]byte("+")),
							godrink.String([]byte("-")),
						),
						godrink.Next(
							godrink.MultiSpace0,
							term,
						),
					),
				),
			),
		),
		func(parsed godrink.Pair[float64, []godrink.Pair[[]byte, godrink.Pair[[]byte, godrink.Pair[[]byte, float64]]]]) float64 {
			val := parsed.Prev
			for _, v := range parsed.Next {
				if v.Next.Prev[0] == '+' {
					val += v.Next.Next.Next
				} else if v.Next.Prev[0] == '-' {
					val -= v.Next.Next.Next
				}
			}
			return val
		},
	)(input)
}
