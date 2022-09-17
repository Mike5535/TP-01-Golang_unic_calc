package cut_string

import (
    "testing"
    "github.com/stretchr/testify/require"
	"github.com/mike5535/uniq/uniq_types"
)

//проверка на удаление пробела вначале
func TestCutSymbolsDelSpace(t *testing.T){

    got := CutSymbols(" abc", 1)
    want := "bc"

	require.Equal(t, got, want, "Cut last")

}

//проверка на пустую строку если число больше длинны строки
func TestCutSymbolsBigNum(t *testing.T){

    got := CutSymbols("abc df", 6)
    want := ""

    require.Equal(t, got, want, "Cut all")
}

//проверка на пустую строку если число больше длинны строки
func TestCutStringBigNum(t *testing.T){

    testParam := uniq_types.OptFields{false,2,0}

    got := CutString("abc df", testParam)
    want := ""

    require.Equal(t, got, want, "Cut all")
}

//проверка на регистр
func TestCutStringRegister(t *testing.T){

    testParam := uniq_types.OptFields{true,1,0}

    got := CutString("Anz Dfa BB", testParam)
    want := "dfa bb"

    require.Equal(t, got, want, "lowercase")
}
