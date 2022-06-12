package roman

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		number int
		want   string
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
		t.Run(fmt.Sprintf("%d got converted to %s", test.number, test.want), func(t *testing.T) {
			got := ConvertToRoman(test.number)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}

}
