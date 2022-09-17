package calc

import (
	"strconv"
	"strings"
)

func CalcTwoNum(num1 int, num2 int, operation string) int {
	res := 0
	if operation == "*" {
		res = num1 * num2
	}
	if operation == "/" {
		res = num1 / num2
	}
	if operation == "-" {
		res = num1 - num2
	}
	if operation == "+" {
		res = num1 + num2
	}
	return res
}

//делает из строки набор аргументов и операций
func ParseSimpleExpr(input string) []string {
	charsIn := strings.Split(input, "")
	if charsIn[0] == "(" {
		charsIn = charsIn[1 : len(charsIn)-1]
	}
	chars := []string{}

	helpFunc := func(pos int) (string, int) {
		result := charsIn[pos]
		pos++

		for _, err := strconv.Atoi(charsIn[pos]); err == nil && pos < len(charsIn); {
			result += charsIn[pos]
			pos++
			if pos < len(charsIn) {
				_, err = strconv.Atoi(charsIn[pos])
			}
		}

		return result, pos
	}

	tempChar := ""
	for i := 0; i < len(charsIn); {
		tempChar = ""

		if i+1 < len(charsIn) {
			_, err1 := strconv.Atoi(charsIn[i+1])
			_, err2 := strconv.Atoi(charsIn[i])
			if err1 == nil && err2 == nil {
				tempChar, i = helpFunc(i)
			}
		}

		if i+1 < len(charsIn) && charsIn[i+1] == "." {
			tempChar = charsIn[i] + charsIn[i+1]
			i += 2
			_, err1 := strconv.Atoi(charsIn[i+1])
			_, err2 := strconv.Atoi(charsIn[i])
			if err1 == nil && err2 == nil {
				bufChar := ""
				bufChar, i = helpFunc(i)
				tempChar += bufChar
			}
		}

		if i < len(charsIn) && tempChar == "" {
			tempChar += charsIn[i]
			i += 1
		}

		chars = append(chars, tempChar)
	}

	return chars
}

//вычисляет выражение без скобок
func CalcSimpleExpr(chars []string) int {
	result := 0
	temp1 := 0
	temp2 := 0
	operation := ""

	temp1, _ = strconv.Atoi(chars[0])

	for i := 0; i < len(chars)-1; {
		if i != 0 {
			temp1 = result
		}

		temp2, _ = strconv.Atoi(chars[i+2])
		operation = chars[i+1]
		if (i+4) < len(chars) && (chars[i+3] == "*" || chars[i+3] == "/") {
			temp3, _ := strconv.Atoi(chars[i+4])
			temp2 = CalcTwoNum(temp2, temp3, chars[i+3])
			i += 4
		} else {
			i += 2
		}

		result = CalcTwoNum(temp1, temp2, operation)
	}

	return result
}

func MakeSubstring(start int, end int, input []string) string {
	var resStr string

	for i := start; i < end+1; i++ {
		resStr += input[i]
	}

	return resStr
}

//выполняет все операции в скобках
func ParseExpr(in string) string {
	input := strings.Split(in, "")
	result := in

	for i := 0; i < len(input); i++ {

		if input[i] == ")" {

			j := i
			for ; input[j] != "("; j-- {
			}

			substring := MakeSubstring(j, i, input)

			resNum := CalcSimpleExpr(ParseSimpleExpr(substring))

			before, after, _ := strings.Cut(in, substring)

			result = ParseExpr(before + strconv.Itoa(resNum) + after)

			break
		}

	}

	return result
}
