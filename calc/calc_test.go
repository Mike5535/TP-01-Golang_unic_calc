package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcTwoNumPlus(t *testing.T) {

	require.Equal(t, CalcTwoNum(1, 2, "+"), 3, "Make plus")

}

func TestCalcTwoNumMinus(t *testing.T) {

	require.Equal(t, CalcTwoNum(5, 2, "-"), 3, "Make subtraction")

}

func TestCalcTwoNumDivide(t *testing.T) {

	require.Equal(t, CalcTwoNum(6, 3, "/"), 2, "Make divide")

}

func TestCalcTwoNumMultiplication(t *testing.T) {

	require.Equal(t, CalcTwoNum(5, 2, "*"), 10, "Make multiplication")

}

func TestParseSimpleExpr(t *testing.T) {

	require.Equal(t, ParseSimpleExpr("1+2*4"), []string{"1","+","2","*","4"}, "делает из строки набор аргументов и операций")

}

func TestParseSimpleExprCutParentheses(t *testing.T) {

	require.Equal(t, ParseSimpleExpr("(4*5+6)"), []string{"4","*","5","+","6"}, "обрезает скобки")

}

func TestCalcSimpleExpr(t *testing.T) {

	require.Equal(t, CalcSimpleExpr([]string{"4","*","5","+","6"}), 26, "вычисляет выражение без скобок")

}

func TestMakeSubstring(t *testing.T) {

    input := []string{"4","*","(","5","+","6",")"}
	require.Equal(t, MakeSubstring(2,6,input), "(5+6)", "создаёт строку")

}

func TestParseExpr(t *testing.T) {

	require.Equal(t, ParseExpr("((3+5)*6+1)+2*4"), "49+2*4", "упрощает выражение")

}