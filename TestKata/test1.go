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
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanNumerals[rune(s[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result
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
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var result strings.Builder
	for i := len(numbers) - 1; i >= 0; i-- {
		for s >= numbers[i] {
			result.WriteString(romans[i])
			s -= numbers[i]
		}
	}

	return result.String()
}
