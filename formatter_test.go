package money

import (
	"testing"
)

func TestFormatter_Format(t *testing.T) {
	tcs := []struct {
		fraction int
		decimal  string
		thousand string
		grapheme string
		template string
		amount   int
		expected string
	}{
		{2, ".", ",", "$", "%s $", 0, "0.00 $"},
		{2, ".", ",", "$", "%s $", 1, "0.01 $"},
		{2, ".", ",", "$", "%s $", 12, "0.12 $"},
		{2, ".", ",", "$", "%s $", 123, "1.23 $"},
		{2, ".", ",", "$", "%s $", 1234, "12.34 $"},
		{2, ".", ",", "$", "%s $", 12345, "123.45 $"},
		{2, ".", ",", "$", "%s $", 123456, "1,234.56 $"},
		{2, ".", ",", "$", "%s $", 1234567, "12,345.67 $"},
		{2, ".", ",", "$", "%s $", 12345678, "123,456.78 $"},
		{2, ".", ",", "$", "%s $", 123456789, "1,234,567.89 $"},

		{2, ".", ",", "$", "%s $", -1, "-0.01 $"},
		{2, ".", ",", "$", "%s $", -12, "-0.12 $"},
		{2, ".", ",", "$", "%s $", -123, "-1.23 $"},
		{2, ".", ",", "$", "%s $", -1234, "-12.34 $"},
		{2, ".", ",", "$", "%s $", -12345, "-123.45 $"},
		{2, ".", ",", "$", "%s $", -123456, "-1,234.56 $"},
		{2, ".", ",", "$", "%s $", -1234567, "-12,345.67 $"},
		{2, ".", ",", "$", "%s $", -12345678, "-123,456.78 $"},
		{2, ".", ",", "$", "%s $", -123456789, "-1,234,567.89 $"},

		{3, ".", "", "$", "%s $", 1, "0.001 $"},
		{3, ".", "", "$", "%s $", 12, "0.012 $"},
		{3, ".", "", "$", "%s $", 123, "0.123 $"},
		{3, ".", "", "$", "%s $", 1234, "1.234 $"},
		{3, ".", "", "$", "%s $", 12345, "12.345 $"},
		{3, ".", "", "$", "%s $", 123456, "123.456 $"},
		{3, ".", "", "$", "%s $", 1234567, "1234.567 $"},
		{3, ".", "", "$", "%s $", 12345678, "12345.678 $"},
		{3, ".", "", "$", "%s $", 123456789, "123456.789 $"},

		{2, ".", ",", "£", "$%s", 1, "£0.01"},
		{2, ".", ",", "£", "$%s", 12, "£0.12"},
		{2, ".", ",", "£", "$%s", 123, "£1.23"},
		{2, ".", ",", "£", "$%s", 1234, "£12.34"},
		{2, ".", ",", "£", "$%s", 12345, "£123.45"},
		{2, ".", ",", "£", "$%s", 123456, "£1,234.56"},
		{2, ".", ",", "£", "$%s", 1234567, "£12,345.67"},
		{2, ".", ",", "£", "$%s", 12345678, "£123,456.78"},
		{2, ".", ",", "£", "$%s", 123456789, "£1,234,567.89"},
	}

	for _, tc := range tcs {
		formatter := NewFormatter(tc.fraction, tc.decimal, tc.thousand, tc.grapheme, tc.template)
		r := formatter.Format(tc.amount)

		if r != tc.expected {
			t.Errorf("Expected %d formatted to be %s got %s", tc.amount, tc.expected, r)
		}
	}
}