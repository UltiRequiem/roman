package roman

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
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

func ParseRoman(romanNumeral string) (uint16, error) {
	var result uint16 = 0

	for _, numeral := range RomanNumerals {
		counter := 0

		for strings.HasPrefix(romanNumeral, numeral.Symbol) {
			if counter > 2 {
				return 0, errors.New("Invalid Roman numeral.")
			}

			result += numeral.Value

			romanNumeral = romanNumeral[len(numeral.Symbol):]

			counter++
		}
	}

	return result, nil
}

func ConvertToRoman(number uint16) (string, error) {
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

func InRomanLimits(number uint16) error {
	if number > 3999 {
		return errors.New("The biggest number we can form in Roman numerals is MMMCMXCIX (3999).")
	}

	if number < 1 {
		return errors.New("There is no concept of 0 or negatives in Roman numerals.")
	}

	return nil
}
