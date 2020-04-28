package numerals

import (
	"fmt"
	"testing"
)

func TestConverToRomans(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{47, "XLVII"},
		{2020, "MMXX"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("convert %d to %s", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRomans(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %s, want %s", got, test.Roman)
			}
		})
		t.Run(fmt.Sprintf("convert %s to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}
