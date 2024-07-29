package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a1, b1, sign, num string //объявляем переменные
	var romanToArabic = map[string]int{
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
	//Принимаем строку
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	weight := scanner.Text()

	//Ищем числа в принятой строке
	for i := 0; i < len(weight); i++ {
		if weight[i] != '+' && weight[i] != '-' && weight[i] != '/' && weight[i] != '*' {
			num += string(weight[i])
		} else {
			if a1 == "" {
				a1 = num
				num = ""
				sign = string(weight[i])
			}
		}
	}
	b1 = num

	//Форматируем строки в число
	if a1 == "1" || a1 == "2" || a1 == "3" || a1 == "4" || a1 == "5" || a1 == "6" || a1 == "7" || a1 == "8" || a1 == "9" || a1 == "10" && b1 == "1" || b1 == "2" || b1 == "3" || b1 == "4" || b1 == "5" || b1 == "6" || b1 == "7" || b1 == "8" || b1 == "9" || b1 == "10" {
		a, _ := strconv.Atoi(a1)
		b, _ := strconv.Atoi(b1)

		//Проверяем числа на условия
		switch {
		case a < 1 || a > 10 || a%10 == 0:
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми")

		case b < 1 || b > 10 || b%10 == 0:
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми")

		}

		switch {
		case sign == "+":
			fmt.Printf("%d + %d = %d\n", a, b, a+b)
		case sign == "-":
			fmt.Printf("%d - %d = %d\n", a, b, a-b)
		case sign == "*":
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
		case sign == "/":
			if a%b != 0 {
				panic("Результатом могут быть целые числа")
			}
			fmt.Printf("%d / %d = %d\n", a, b, a/b)
		}
	} else {
		a := romanToArabic[a1]
		b := romanToArabic[b1]

		switch {
		case a < 1 || a > 10 || a%10 == 0:
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми")

		case b < 1 || b > 10 || b%10 == 0:
			panic("Должны быть числа от 1 до 10 включительно. Числа должны быть целыми")
		}
		//Выполняем действия исходя из условий

		switch {
		case sign == "+":
			fmt.Printf("%d + %d = %d\n", a, b, a+b)
		case sign == "-":
			fmt.Printf("%d - %d = %d\n", a, b, a-b)
		case sign == "*":
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
		case sign == "/":
			if a%b != 0 {
				panic("Результатом могут быть целые числа")
			}
			fmt.Printf("%d / %d = %d\n", a, b, a/b)
		}
	}
}
