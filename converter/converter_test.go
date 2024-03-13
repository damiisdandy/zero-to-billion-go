package converter

import (
	"fmt"
	"testing"
)

type testTree []struct {
	num  int
	want string
}

func TestConvertUnit(t *testing.T) {
	tests := testTree{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, ""},
		{11, "one"},
		{102, "two"},
		{47657, "seven"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Get unit for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertUnit(tt.num); got != tt.want {
				t.Errorf("ConvertUnit() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertTens(t *testing.T) {
	tests := testTree{
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{30, "thirty"},
		{40, "forty"},
		{50, "fifty"},
		{60, "sixty"},
		{70, "seventy"},
		{80, "eighty"},
		{90, "ninety"},
		{100, ""},
		{101, ""},
		{113, "thirteen"},
		{120, "twenty"},
		{2374756375, "seventy"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Get tens for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertTens(tt.num); got != tt.want {
				t.Errorf("ConvertTens() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertHundred(t *testing.T) {
	tests := testTree{
		{100, "one hundred"},
		{123, "one hundred and twenty three"},
		{301, "three hundred and one"},
		{47574854, "eight hundred and fifty four"},
		{49, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get hundreds for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertHundred(tt.num); got != tt.want {
				t.Errorf("ConvertHundred() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertHundredAndBelow(t *testing.T) {
	tests := testTree{
		{0, "zero"},
		{17, "seventeen"},
		{18, "eighteen"},
		{83, "eighty three"},
		{100, "one hundred"},
		{123, "one hundred and twenty three"},
		{301, "three hundred and one"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get hundreds and below for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertHundredAndBelow(tt.num, false); got != tt.want {
				t.Errorf("ConvertHundredAndBelow() = %q, want %q", got, tt.want)
			}
		})
	}

	t.Run("With zeros hidden", func(t *testing.T) {
		if got := ConvertHundredAndBelow(0, true); got != "" {
			t.Errorf("ConvertHundredAndBelow() = %q, want %q", got, "")
		}
	})
}

func TestConvertThousand(t *testing.T) {
	tests := testTree{
		{1000, "one thousand"},
		{1234, "one thousand two hundred and thirty four"},
		{3010, "three thousand and ten"},
		{1999, "one thousand nine hundred and ninety nine"},
		{100000, "one hundred thousand"},
		{123456, "one hundred and twenty three thousand four hundred and fifty six"},
		{574854, "five hundred and seventy four thousand eight hundred and fifty four"},
		{49, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get thousands for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertThousand(tt.num); got != tt.want {
				t.Errorf("ConvertThousand() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertMillion(t *testing.T) {
	tests := testTree{
		{1_000_000, "one million"},
		{1_000_001, "one million and one"},
		{1_000_101, "one million one hundred and one"},
		{1_234_567, "one million two hundred and thirty four thousand five hundred and sixty seven"},
		{3_010_000, "three million ten thousand"},
		{1_999_999, "one million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000, "one hundred million"},
		{123_456_789, "one hundred and twenty three million four hundred and fifty six thousand seven hundred and eighty nine"},
		{999_999_999, "nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get millions for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertMillion(tt.num); got != tt.want {
				t.Errorf("ConvertMillion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertBillion(t *testing.T) {
	tests := testTree{
		{1_000_000_000, "one billion"},
		{1_000_000_001, "one billion and one"},
		{1_000_100_001, "one billion one hundred thousand and one"},
		{1_234_567_890, "one billion two hundred and thirty four million five hundred and sixty seven thousand eight hundred and ninety"},
		{3_010_000_000, "three billion ten million"},
		{1_999_999_999, "one billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000_000, "one hundred billion"},
		{123_456_789_012, "one hundred and twenty three billion four hundred and fifty six million seven hundred and eighty nine thousand and twelve"},
		{999_999_999_999, "nine hundred and ninety nine billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get billions for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertBillion(tt.num); got != tt.want {
				t.Errorf("ConvertBillion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConvertTrillion(t *testing.T) {
	tests := testTree{
		{1_000_000_000_000, "one trillion"},
		{1_000_000_000_023, "one trillion and twenty three"},
		{1_000_000_020_023, "one trillion twenty thousand and twenty three"},
		{1_000_000_000_001, "one trillion and one"},
		{1_000_100_000_001, "one trillion one hundred million and one"},
		{1_234_567_890_123, "one trillion two hundred and thirty four billion five hundred and sixty seven million eight hundred and ninety thousand one hundred and twenty three"},
		{3_010_000_000_000, "three trillion ten billion"},
		{1_999_999_999_999, "one trillion nine hundred and ninety nine billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000_000_000, "one hundred trillion"},
		{123_456_789_012_345, "one hundred and twenty three trillion four hundred and fifty six billion seven hundred and eighty nine million twelve thousand three hundred and forty five"},
		{999_999_999_999_999, "nine hundred and ninety nine trillion nine hundred and ninety nine billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Get trillions for %d", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := ConvertTrillion(tt.num); got != tt.want {
				t.Errorf("ConvertTrillion() = %q, want %q", got, tt.want)
			}
		})
	}

}

func TestConverter(t *testing.T) {
	tests := testTree{
		{0, "zero"},
		{5, "five"},
		{13, "thirteen"},
		{20, "twenty"},
		{30, "thirty"},
		{34, "thirty four"},
		{99, "ninety nine"},
		{100, "one hundred"},
		{101, "one hundred and one"},
		{123, "one hundred and twenty three"},
		{423, "four hundred and twenty three"},
		{999, "nine hundred and ninety nine"},
		{1000, "one thousand"},
		{1234, "one thousand two hundred and thirty four"},
		{3010, "three thousand and ten"},
		{1_000_101, "one million one hundred and one"},
		{1_234_567, "one million two hundred and thirty four thousand five hundred and sixty seven"},
		{3_010_000, "three million ten thousand"},
		{1_999_999, "one million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000, "one hundred million"},
		{1_234_567_890, "one billion two hundred and thirty four million five hundred and sixty seven thousand eight hundred and ninety"},
		{3_010_000_000, "three billion ten million"},
		{1_999_999_999, "one billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000_000, "one hundred billion"},
		{1_234_567_890_123, "one trillion two hundred and thirty four billion five hundred and sixty seven million eight hundred and ninety thousand one hundred and twenty three"},
		{3_010_000_000_000, "three trillion ten billion"},
		{1_999_999_999_999, "one trillion nine hundred and ninety nine billion nine hundred and ninety nine million nine hundred and ninety nine thousand nine hundred and ninety nine"},
		{100_000_000_000_000, "one hundred trillion"},
		{1_000_000_000_000_000, "one quadrillion"},
		{1_000_000_000_000_001, "Number is out of range"},
		{-1, "Number is out of range"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Convert %d to word", tt.num)
		t.Run(testName, func(t *testing.T) {
			if got := Converter(tt.num); got != tt.want {
				t.Errorf("Converter() = %q, want %q", got, tt.want)
			}
		})
	}
}
