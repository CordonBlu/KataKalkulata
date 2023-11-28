package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите задачу: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	result, err := calculate(input)
	j := strings.Split(input, " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result1 string
	if isRomanNumeral(j[0]) && isRomanNumeral(j[2]) {
		if result > 0 {
			result1 = arabicToRoman(result)
			fmt.Println(result1)
		} else {
			fmt.Println("Ошибка. Ответ может быть только натуральным.")
			return
		}
	} else {
		fmt.Println(result)
		return
	}

}

func calculate(input string) (int, error) {
	// разделяем строку на числа и операцию
	values := strings.Split(input, " ")
	if len(values) != 3 {
		return 0, fmt.Errorf("Неверный формат ввода")
	}
	// проверяем, являются ли числа римскими или арабскими
	isRoman := isRomanNumeral(values[0]) && isRomanNumeral(values[2])
	isArabic := isArabicNumeral(values[0]) && isArabicNumeral(values[2])
	if !isRoman && !isArabic {
		return 0, fmt.Errorf("неверный формат чисел")
	}
	if isRoman && isArabic {
		return 0, fmt.Errorf("числа должны быть одного типа")
	}

	// преобразуем числа в нужный формат
	var a, b int
	if isRoman {
		a = romanToArabic(values[0])
		b = romanToArabic(values[2])
	} else {
		a = arabicToInteger(values[0])
		b = arabicToInteger(values[2])
	}

	// выполняем операцию

	var result int
	switch values[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("неверная операция")
	}
	return result, nil
}

func isRomanNumeral(s string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range romanNumerals {
		if s == numeral {
			return true
		}
	}
	return false
}

func isArabicNumeral(s string) bool {
	arabicNumerals := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, numeral := range arabicNumerals {
		if s == numeral {
			return true
		}
	}
	return false
}

func romanToArabic(s string) int {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanNumerals[s]
}

func arabicToInteger(s string) int {
	arabicNumerals := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	return arabicNumerals[s]
}
func arabicToRoman(s int) string {
	arabicNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return arabicNumerals[s]
}
