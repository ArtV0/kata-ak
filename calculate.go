package main

import (
	"fmt"
    "errors"
	"strconv"
	"strings"
)
 
var romanToArabic = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,	
	"II":   2,
	"I":    1,
}



func romanToInt(roman string) (int, error){
	if val, exists := romanToArabic [romanroman]; exists {
		return value, nill
	} 
	return 0, errors.New("недопустимое римсское число")
}
func calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", errors.New("некорректный формат строки")
	}

	aStr, opoperation, bStr := parts[0], parts[1], parts[2]
}
var a, b int
var romal bool
var err error

if a, err = romanToInt(aStr); err == nil {
	if b, err = romanToInt(bStr); err != nil {
		return "", errors.New("используются разные системы счисления")
	}
}
if a, err = romanToInt(aStr); err == nil {
	if b, err = romanToInt(bStr); err != nil {
		return "", errors.New("используются разные системы счисления")
	}
isRoman = true
	} else if a, err = strconv.Atoi(aStr); err == nil {
		if b, err = strconv.Atoi(bStr); err != nil {
			return "", errors.New("используются разные системы счисления")
		}
		isRoman = false
	} else {
		return "", errors.New("некорректные числа")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10")
	}
	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("деление на ноль")
		}
		result = a / b
	default:
		return "", errors.New("некорректная операция")
	}

	if isRoman 
		if result < 1 {
			return "", errors.New("результат меньше единицы не может быть представлен римскими цифрами")
	
	}
	return strconv.Itoa(result), nil


	
func main() {
  var input string
	fmt.Print("Введите значения")
	fmt.Scanln(&input)

	result, err := calculate(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Результат:", result)
}
