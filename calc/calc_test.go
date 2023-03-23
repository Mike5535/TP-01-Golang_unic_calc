package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcTwoNumPlus(t *testing.T) {
	actual := CalcTwoNum(1, 2, "+")
	expected := 3

	require.Equal(t, actual, expected, "Make plus")
}

func TestCalcTwoNumMinus(t *testing.T) {
	actual := CalcTwoNum(5, 2, "-")
	expected := 3

	require.Equal(t, actual, expected, "Make subtraction")
}

func TestCalcTwoNumDivide(t *testing.T) {
	actual := CalcTwoNum(6, 3, "/")
	expected := 2

	require.Equal(t, actual, expected, "Make divide")
}

func TestCalcTwoNumMultiplication(t *testing.T) {
	actual := CalcTwoNum(5, 2, "*")
	expected := 10

	require.Equal(t, actual, expected, "Make multiplication")
}

func TestParseSimpleExpr(t *testing.T) {
	actual := ParseSimpleExpr("1+2*4")
	expected := []string{"1","+","2","*","4"}

	require.Equal(t, actual, expected, "делает из строки набор аргументов и операций")
}

func TestParseSimpleExprCutParentheses(t *testing.T) {
	actual := ParseSimpleExpr("(4*5+6)")
	expected := []string{"4","*","5","+","6"}

	require.Equal(t, actual, expected, "обрезает скобки")
}

func TestCalcSimpleExpr(t *testing.T) {
	actual := CalcSimpleExpr([]string{"4","*","5","+","6"})
	expected := 26

	require.Equal(t, actual, expected, "вычисляет выражение без скобок")
}

func TestMakeSubstring(t *testing.T) {
    input := []string{"4","*","(","5","+","6",")"}
	actual := MakeSubstring(2,6,input)
	expected := "(5+6)"

	require.Equal(t, actual, expected, "создаёт строку")
}

func TestParseExpr(t *testing.T) {
	actual := ParseExpr("((3+5)*6+1)+2*4")
	expected := "49+2*4"

	require.Equal(t, actual, expected, "упрощает выражение")
}