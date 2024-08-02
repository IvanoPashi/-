package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

func romanToInt(roman string) (int, error) {
	if value, exists := romanToArabic[roman]; exists {
		return value, nil
	}
	return 0, fmt.Errorf("Invalid Roman numeral: %s", roman)
}

func intToRoman(number int) (string, error) {
	if number > 0 && number <= 10 {
		return arabicToRoman[number], nil
	}
	return "", fmt.Errorf("Resulting Roman numeral is out of bounds: %d", number)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите первое число:")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Println("Введите оператор (+, -, *, /):")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	fmt.Println("Введите второе число:")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	var num1, num2 int
	var err error
	var romanInput bool

	if _, exists := romanToArabic[input1]; exists {
		romanInput = true
		num1, err = romanToInt(input1)
		if err != nil {
			panic(err)
		}
		num2, err = romanToInt(input2)
		if err != nil {
			panic(err)
		}
	} else {
		num1, err = strconv.Atoi(input1)
		if err != nil || num1 < 1 || num1 > 10 {
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми.")
		}
		num2, err = strconv.Atoi(input2)
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
			panic("Resulting Roman numeral is less than I")
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
