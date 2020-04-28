package numerals

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(s string) int {
	for _, sym := range r {
		if sym.Symbol == s {
			return sym.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= uint16(numeral.Value) {
			result.WriteString(numeral.Symbol)
			arabic -= uint16(numeral.Value)
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var result int
	for i, r := range roman {
		value := allRomanNumerals.ValueOf(string(r))
		multiplier := getValueMultiplier(i, value, roman)
		result += value * multiplier
	}
	return uint16(result)
}

func getValueMultiplier(i, value int, roman string) int {
	multiplier := 1
	if i+1 < len(roman) {
		next := roman[i+1]
		if allRomanNumerals.ValueOf(string(next)) > value {
			multiplier = -1

		}
	}
	return multiplier
}
