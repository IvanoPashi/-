package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Map to convert Roman numerals to Arabic numerals
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Map to convert Arabic numerals to Roman numerals
var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

// Convert Roman numeral to Arabic numeral
func romanToInt(roman string) (int, error) {
	if value, exists := romanToArabic[roman]; exists {
		return value, nil
	}
	return 0, fmt.Errorf("Invalid Roman numeral: %s", roman)
}

// Convert Arabic numeral to Roman numeral
func intToRoman(number int) (string, error) {
	if number > 0 {
		result := ""
		for number >= 10 {
			result += "X"
			number -= 10
		}
		if number > 0 {
			result += arabicToRoman[number]
		}
		return result, nil
	}
	return "", fmt.Errorf("Resulting Roman numeral is out of bounds: %d", number)
}

func main() {
	var input string
	fmt.Println("Введите выражение (например, 3 + 4 или IV * VI):")
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("Failed to read input")
	}

	// Split the input by space
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Input should be in the format: number operator number")
	}

	num1Str, operator, num2Str := parts[0], parts[1], parts[2]

	var num1, num2 int
	var romanInput bool

	// Determine if the input is Roman or Arabic
	if _, exists := romanToArabic[num1Str]; exists {
		romanInput = true
		num1, err = romanToInt(num1Str)
		if err != nil {
			panic(err)
		}
		num2, err = romanToInt(num2Str)
		if err != nil {
			panic(err)
		}
	} else {
		num1, err = strconv.Atoi(num1Str)
		if err != nil || num1 < 1 || num1 > 10 {
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми.")
		}
		num2, err = strconv.Atoi(num2Str)
		if err != nil || num2 < 1 || num2 > 10 {
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми.")
		}
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Division by zero")
		}
		result = num1 / num2
	default:
		panic("Invalid operator")
	}

	if romanInput {
		if result < 1 {
			panic("Результат вычисления римских чисел меньше I")
		}
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Результат: %s\n", romanResult)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
