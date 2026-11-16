package roman

import (
	"fmt"
	"testing"
)

var cases = []struct {
	value   uint16
	numeral string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{40, "XL"},
	{47, "XLVII"},
	{50, "L"},
	{798, "DCCXCVIII"},
	{1006, "MVI"},
	{1984, "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d got converted to %s", test.value, test.numeral), func(t *testing.T) {
			got, errrorConverting := ConvertToRoman(test.value)

			if errrorConverting != nil {
				t.Error(errrorConverting)
			}

			if got != test.numeral {
				t.Errorf("got %v, expected %v", got, test.numeral)
			}
		})
	}

	t.Run("Number lower than 1 returns an error", func(t *testing.T) {
		_, error := ConvertToRoman(0)

		if error == nil {
			t.Error("Expected an error, didn't get one.")
		}
	})

	t.Run("Number greater than 3999 returns an error", func(t *testing.T) {
		_, error := ConvertToRoman(5000)

		if error == nil {
			t.Error("Expected an error, didn't get one.")
		}
	})

}

func TestParseRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s got converted to %d.", test.numeral, test.value), func(t *testing.T) {
			got, errrorConverting := ParseRoman(test.numeral)

			if errrorConverting != nil {
				t.Error(errrorConverting)
			}

			if got != test.value {
				t.Errorf("Got %d, expected %v.", got, test.value)
			}
		})
	}

	t.Run("Invalid Roman numerals returns an error.", func(t *testing.T) {
		_, errorParsing := ParseRoman("XXXX")

		if errorParsing == nil {
			t.Error("Expected an error, didn't get one.")
		}
	})
}

func TestParseRomanRejectsNonCanonical(t *testing.T) {
	for _, input := range []string{"IIX", "VX", "IL", "IVI", "XCX", "VV", "CMCM", "MCMXCIV ", "ABC", "iv", "", "é", "MMMM"} {
		if _, err := ParseRoman(input); err == nil {
			t.Errorf("expected an error parsing %q", input)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for n := uint16(1); n <= 3999; n++ {
		numeral, err := ConvertToRoman(n)
		if err != nil {
			t.Fatal(err)
		}

		got, err := ParseRoman(numeral)
		if err != nil || got != n {
			t.Fatalf("%d -> %s -> %d (%v)", n, numeral, got, err)
		}
	}
}

func BenchmarkConvertToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := uint16(1); n <= 3999; n++ {
			_, _ = ConvertToRoman(n)
		}
	}
}

func BenchmarkParseRoman(b *testing.B) {
	numerals := make([]string, 0, 3999)
	for n := uint16(1); n <= 3999; n++ {
		s, _ := ConvertToRoman(n)
		numerals = append(numerals, s)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, s := range numerals {
			_, _ = ParseRoman(s)
		}
	}
}
