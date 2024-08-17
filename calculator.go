package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3,
	"IV": 4, "V": 5, "VI": 6, "VII": 7,
	"VIII": 8, "IX": 9, "X": 10,
}
var arabRoman = []struct {
	value  int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"},
	{40, "XL"}, {10, "X"}, {9, "IX"},
	{5, "V"}, {4, "IV"}, {1, "I"},
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неверное выражение")
	}
}
func romanToInt(roman string) int {
	value, exists := romanNumerals[roman]
	if !exists {
		panic("Введите римское число")
	}
	return value
}
func intToRoman(num int) string {
	if num <= 0 {
		panic("Римские числа не могут быть меньше единицы")
	}
	result := ""
	for _, x := range arabRoman {
		for num >= x.value {
			result += x.symbol
			num -= x.value
		}
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите арифметическую операцию:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		panic("Неверный формат ввода")
	}

	aStr, operator, bStr := tokens[0], tokens[1], tokens[2]

	isRoman := false
	if _, exits := romanNumerals[aStr]; exits {
		isRoman = true
	}

	var a, b int
	if isRoman == true {
		a = romanToInt(aStr)
		b = romanToInt(bStr)
	} else {
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic("Недопустимое арабское число")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic("Недопустимое арабское число")
		}
	}
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10 включительно")
	}
	result := calculate(a, b, operator)

	if isRoman {
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}
