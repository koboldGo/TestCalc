package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cases = [19]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var arabicRegex = regexp.MustCompile("^(10|[1-9])?$")
var romeRegex = regexp.MustCompile("[IVXivx]")
var operationRegex = regexp.MustCompile("[/+\\-*]")

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		validate(strings.ReplaceAll(text, "\n", ""))
	}
}

func validate(input string) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	operands := operationRegex.Split(input, -1)
	if len(operands) != 2 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	if arabicRegex.Match([]byte(operands[0])) && arabicRegex.Match([]byte(operands[1])) {
		calcArabic(input, operands)
	} else if romeRegex.Match([]byte(operands[0])) && romeRegex.Match([]byte(operands[1])) {
		calcRome(input, operands)
	} else {
		panic("Выражение не соответствует требованиям")
	}
}

func calcArabic(input string, operands []string) {
	first, _ := strconv.Atoi(operands[0])
	second, _ := strconv.Atoi(operands[1])
	switch {
	case strings.Contains(input, "+"):
		fmt.Println(first + second)
	case strings.Contains(input, "-"):
		fmt.Println(first - second)
	case strings.Contains(input, "/"):
		fmt.Println(first / second)
	case strings.Contains(input, "*"):
		fmt.Println(first * second)
	}
}

func calcRome(input string, operands []string) {
	first := convertToArabic(operands[0])
	second := convertToArabic(operands[1])
	switch {
	case strings.Contains(input, "+"):
		convertToRome(first + second)
	case strings.Contains(input, "-"):
		convertToRome(first - second)
	case strings.Contains(input, "/"):
		convertToRome(first / second)
	case strings.Contains(input, "*"):
		convertToRome(first * second)
	}
}

func convertToArabic(input string) int {
	switch {
	case input == "I":
		{
			return 1
		}
	case input == "II":
		{
			return 2
		}
	case input == "III":
		{
			return 3
		}
	case input == "IV":
		{
			return 4
		}
	case input == "V":
		{
			return 5
		}
	case input == "VI":
		{
			return 6
		}
	case input == "VII":
		{
			return 7
		}
	case input == "VIII":
		{
			return 8
		}
	case input == "IX":
		{
			return 9
		}
	case input == "X":
		{
			return 10
		}
	}
	return 0
}

func convertToRome(input int) {
	if input < 0 {
		panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
	}
	var buffer bytes.Buffer
	for _, element := range cases {
		if input >= element {
			switch {
			case element == 1:
				buffer.WriteString("I")
			case element == 2:
				buffer.WriteString("II")
			case element == 3:
				buffer.WriteString("III")
			case element == 4:
				buffer.WriteString("IV")
			case element == 5:
				buffer.WriteString("V")
			case element == 6:
				buffer.WriteString("VI")
			case element == 7:
				buffer.WriteString("VII")
			case element == 8:
				buffer.WriteString("VIII")
			case element == 9:
				buffer.WriteString("IX")
			case element == 10:
				buffer.WriteString("X")
			case element == 20:
				buffer.WriteString("XX")
			case element == 30:
				buffer.WriteString("XXX")
			case element == 40:
				buffer.WriteString("XL")
			case element == 50:
				buffer.WriteString("L")
			case element == 60:
				buffer.WriteString("LX")
			case element == 70:
				buffer.WriteString("LXX")
			case element == 80:
				buffer.WriteString("LXXX")
			case element == 90:
				buffer.WriteString("XC")
			case element == 100:
				buffer.WriteString("C")
			}
			input = input - element
		}
	}
	fmt.Println(buffer.String())
}
