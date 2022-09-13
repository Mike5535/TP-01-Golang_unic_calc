package uniq

import (
    "testing"
    "container/list"
)

//проверка на удаление пробела вначале
func TestCutSymbolsDelSpace(t *testing.T){

    got := CutSymbols(" abc", 1)
    want := "bc"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

//проверка на пустую строку если число больше длинны строки
func TestCutSymbolsBigNum(t *testing.T){

    got := CutSymbols("abc df", 6)
    want := ""

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

//проверка на пустую строку если число больше длинны строки
func TestCutStringBigNum(t *testing.T){
    /*type OptFields struct {
        AnyRegister bool
        NumFields   int
        NumChars    int
    }*/

    testParam := OptFields{false,2,0}

    got := CutString("abc df", testParam)
    want := ""

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

//проверка на регистр
func TestCutStringRegister(t *testing.T){
    /*type OptFields struct {
        AnyRegister bool
        NumFields   int
        NumChars    int
    }*/

    testParam := OptFields{true,1,0}

    got := CutString("Anz Dfa BB", testParam)
    want := "dfa bb"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

// проверка подсчёта
func TestUniqleCounting(t *testing.T){
    /*type Options struct {
	Counting     bool
	Repeat       bool
	Uniq         bool
	FieldOptions *OptFields
    }*/
    
    testParam := Options{true,false,false,&OptFields{false,0,0}}
    var input list.List
    for i := 0; i < 3; i++ {
    input.PushBack("I love music.")
    }
    input.PushBack(" ")
    for i := 0; i < 2; i++ {
        input.PushBack("I love music of Kartik.")
    }

    got := Uniqle(testParam, input)

    var want list.List
    want.PushBack(Pair{"I love music.",3})
    want.PushBack(Pair{" ",1})
    want.PushBack(Pair{"I love music of Kartik.",2})

    w := want.Front()
    for e := got.Front(); e != nil; e = e.Next() {
        if e.Value.(Pair).numb != w.Value.(Pair).numb {
            t.Errorf("got %d, wanted %d", e.Value.(Pair).numb, w.Value.(Pair).numb)
        }
        if w.Next() != nil {
            w = w.Next()
        }
    }
}
