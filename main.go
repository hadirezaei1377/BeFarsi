package main

import (
	"fmt"
	"strings"
)

var (
	persianNumbers = []string{"", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"}
	tens           = []string{"", "ده", "بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"}
	teens          = []string{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هجده", "نوزده"}
	hundreds       = []string{"", "یکصد", "دویست", "سیصد", "چهارصد", "پانصد", "ششصد", "هفتصد", "هشتصد", "نهصد"}
	orders         = []string{"", "هزار", "میلیون", "میلیارد", "تریلیون"}
)

func convertThreeDigits(number int) string {
	if number == 0 {
		return ""
	}

	hundred := number / 100
	remainder := number % 100
	tensPlace := remainder / 10
	onesPlace := remainder % 10
	words := []string{}

	if hundred > 0 {
		words = append(words, hundreds[hundred])
		if remainder > 0 {
			words = append(words, "و")
		}
	}

	if remainder > 0 {
		if remainder < 10 {
			words = append(words, persianNumbers[onesPlace])
		} else if remainder >= 10 && remainder < 20 {
			words = append(words, teens[onesPlace])
		} else {
			if onesPlace > 0 {
				words = append(words, tens[tensPlace]+" و "+persianNumbers[onesPlace])
			} else {
				words = append(words, tens[tensPlace])
			}
		}
	}

	return strings.Join(words, " ")
}

func numberToWords(number int64) string {
	if number == 0 {
		return "صفر"
	}

	parts := []string{}
	for order := 0; number > 0; order++ {
		threeDigits := int(number % 1000)
		if threeDigits > 0 {
			part := convertThreeDigits(threeDigits)
			if order > 0 {
				part += " " + orders[order]
			}
			parts = append([]string{part}, parts...)
		}
		number /= 1000
	}

	return strings.Join(parts, " و ")
}

func main() {
	var number int64
	println("Enter a number: ")
	fmt.Scan(&number)
	fmt.Println(numberToWords(number))
}
