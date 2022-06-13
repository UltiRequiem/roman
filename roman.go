package roman

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(number int) (string, error) {
	ableToConvert := InRomanLimits(number)

	if ableToConvert != nil {
		return "", ableToConvert
	}

	var result strings.Builder

	for _, numeral := range RomanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return result.String(), nil
}

func InRomanLimits(number int) error {
	if number > 3999 {
		return errors.New("The biggest number we can form in Roman numerals is MMMCMXCIX (3999).")
	}

	if number < 1 {
		return errors.New("There is no concept of 0 or negatives in Roman numerals.")
	}

	return nil
}
