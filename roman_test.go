package roman

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		given    int
		expected string
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

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d got converted to %s", test.given, test.expected), func(t *testing.T) {
			got, errrorConverting := ConvertToRoman(test.given)

			if errrorConverting != nil {
				t.Error(errrorConverting)
			}

			if got != test.expected {
				t.Errorf("got %v, exected %v", got, test.expected)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	t.Run("Number lower than 1 returns an error", func(t *testing.T) {
		_, error := ConvertToRoman(-3)

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
