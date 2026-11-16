package roman

import "errors"

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

var (
	ones      = [...]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = [...]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds  = [...]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = [...]string{"", "M", "MM", "MMM"}
)

// romanTable holds every Roman numeral from 1 to 3999, indexed by the number
// itself, so conversion is a single array lookup.
var romanTable = func() (table [4000]string) {
	for n := 1; n < len(table); n++ {
		table[n] = thousands[n/1000] + hundreds[n/100%10] + tens[n/10%10] + ones[n%10]
	}

	return
}()

// charValues maps a Roman numeral character to its value (0 if invalid).
var charValues = [256]uint16{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var errInvalidNumeral = errors.New("Invalid Roman numeral.")

func ParseRoman(romanNumeral string) (uint16, error) {
	var result, previous uint16

	// Walk right to left: a smaller value before a bigger one is subtracted.
	for i := len(romanNumeral) - 1; i >= 0; i-- {
		value := charValues[romanNumeral[i]]

		if value == 0 {
			return 0, errInvalidNumeral
		}

		if value < previous {
			result -= value
		} else {
			result += value
		}

		previous = value
	}

	// Every number has exactly one canonical numeral, so anything that does
	// not round-trip (IIX, VX, XXXX, CMCM, the empty string...) is invalid.
	if result == 0 || result > 3999 || romanTable[result] != romanNumeral {
		return 0, errInvalidNumeral
	}

	return result, nil
}

func ConvertToRoman(number uint16) (string, error) {
	if err := InRomanLimits(number); err != nil {
		return "", err
	}

	return romanTable[number], nil
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
