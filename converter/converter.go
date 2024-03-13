package converter

import (
	"fmt"
	"strings"
)

const (
	UNIT_UPPER_LIMIT     = 10
	TENS_UPPER_LIMIT     = 100
	HUNDRED_UPPER_LIMIT  = 1000
	THOUSAND_UPPER_LIMIT = 1_000_000
	MILLION_UPPER_LIMIT  = 1_000_000_000
	BILLION_UPPER_LIMIT  = 1_000_000_000_000
	MAX_UPPER_LIMIT      = 1_000_000_000_000_000
	FIXED_TEXT_RANGE     = 20 // 20 and below
)

func ConvertUnit(num int) string {
	if num == 0 {
		return "zero"
	}
	unitMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	unit := num % UNIT_UPPER_LIMIT
	return unitMap[unit]
}

func ConvertTens(num int) string {
	tensMap := map[int]string{
		0:  "",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
	tens := num % TENS_UPPER_LIMIT

	val, ok := tensMap[tens]
	if ok {
		return val
	} else {
		// e.g 34 - 4 = 30
		unit := tens % UNIT_UPPER_LIMIT
		tens = tens - unit
		return tensMap[tens]

	}
}

func ConvertHundred(num int) (output string) {
	hundred := num / TENS_UPPER_LIMIT
	remainder := num % TENS_UPPER_LIMIT
	if hundred > 0 {
		separator := ""
		if remainder > 0 {
			// e.g 101 - one hundred and one
			separator = " and"
		}
		output = fmt.Sprintf("%s hundred%s", ConvertUnit(hundred), separator)
		outputEnd := fmt.Sprintf("%s %s", ConvertTens(num), ConvertUnit(num))
		// // removing the trailing space when tens is zero
		// // e.g 101 - one hundred and  one
		// // becoomes one hundred and one
		output = fmt.Sprintf("%s %s", output, strings.TrimSpace(outputEnd))
	}
	return strings.TrimSpace(output)
}

// ConvertHundredAndBelow converts numbers from 100 and below
// the conversion for hundred needs to be extracted since it will be used for numbers like
// 102,000,  234,444,000, etc./
func ConvertHundredAndBelow(num int, hideZero bool) (output string) {
	if num == 0 && hideZero {
		return ""
	}
	if num < FIXED_TEXT_RANGE {
		if num < UNIT_UPPER_LIMIT {
			return ConvertUnit(num)
		}
		return ConvertTens(num)
	} else if num < TENS_UPPER_LIMIT {
		output = fmt.Sprintf("%s %s", ConvertTens(num), ConvertUnit(num))
	} else if num < HUNDRED_UPPER_LIMIT {
		output = ConvertHundred(num)
	}
	return strings.TrimSpace(output)
}

func ConvertThousand(num int) (output string) {
	thousand := num / HUNDRED_UPPER_LIMIT
	hundreds := num % HUNDRED_UPPER_LIMIT

	if thousand > 0 {
		separator := ""
		if hundreds < TENS_UPPER_LIMIT && hundreds > 0 {
			separator = " and"
		}
		output = fmt.Sprintf("%s thousand%s %s", ConvertHundredAndBelow(thousand, true), separator, strings.TrimSpace(ConvertHundredAndBelow(hundreds, true)))
	}
	return strings.TrimSpace(output)
}

func ConvertMillion(num int) (output string) {
	million := num / THOUSAND_UPPER_LIMIT
	thousands := num % THOUSAND_UPPER_LIMIT

	if million > 0 {
		separator := ""
		if thousands < TENS_UPPER_LIMIT && thousands > 0 {
			separator = " and"
		}
		thousandsValue := strings.TrimSpace(ConvertThousand(thousands))
		if thousands < HUNDRED_UPPER_LIMIT {
			thousandsValue = strings.TrimSpace(ConvertHundredAndBelow(thousands, true))
		}
		output = fmt.Sprintf("%s million%s %s", ConvertHundredAndBelow(million, true), separator, thousandsValue)
	}
	return strings.TrimSpace(output)
}

func ConvertBillion(num int) (output string) {
	billion := num / MILLION_UPPER_LIMIT
	millions := num % MILLION_UPPER_LIMIT

	if billion > 0 {
		separator := ""
		if millions < TENS_UPPER_LIMIT && millions > 0 {
			separator = " and"
		}
		millionsValue := strings.TrimSpace(ConvertMillion(millions))
		if millions < HUNDRED_UPPER_LIMIT {
			millionsValue = strings.TrimSpace(ConvertHundredAndBelow(millions, true))
		} else if millions < THOUSAND_UPPER_LIMIT {
			millionsValue = strings.TrimSpace(ConvertThousand(millions))
		}
		output = fmt.Sprintf("%s billion%s %s", ConvertHundredAndBelow(billion, true), separator, millionsValue)
	}
	return strings.TrimSpace(output)
}

func ConvertTrillion(num int) (output string) {
	trillion := num / BILLION_UPPER_LIMIT
	billions := num % BILLION_UPPER_LIMIT

	if trillion > 0 {
		separator := ""
		if billions < TENS_UPPER_LIMIT && billions > 0 {
			separator = " and"
		}
		billionsValue := strings.TrimSpace(ConvertBillion(billions))
		if billions < HUNDRED_UPPER_LIMIT {
			billionsValue = strings.TrimSpace(ConvertHundredAndBelow(billions, true))
		} else if billions < THOUSAND_UPPER_LIMIT {
			billionsValue = strings.TrimSpace(ConvertThousand(billions))
		} else if billions < MILLION_UPPER_LIMIT {
			billionsValue = strings.TrimSpace(ConvertMillion(billions))
		}
		output = fmt.Sprintf("%s trillion%s %s", ConvertHundredAndBelow(trillion, true), separator, billionsValue)
	}
	return strings.TrimSpace(output)
}

func Converter(num int) (result string) {
	if num == MAX_UPPER_LIMIT {
		return "one quadrillion"
	}
	if num < HUNDRED_UPPER_LIMIT {
		result = ConvertHundredAndBelow(num, false)
	} else if num < THOUSAND_UPPER_LIMIT {
		result = ConvertThousand(num)
	} else if num < MILLION_UPPER_LIMIT {
		result = ConvertMillion(num)
	} else if num < BILLION_UPPER_LIMIT {
		result = ConvertBillion(num)
	} else if num < MAX_UPPER_LIMIT {
		result = ConvertTrillion(num)
	}
	return strings.TrimSpace(result)
}
