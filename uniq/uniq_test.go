package uniq

import (
    "testing"
    "container/list"
    "github.com/stretchr/testify/require"
    "github.com/mike5535/uniq/uniq_types"
)

// проверка подсчёта
func TestUniqleCounting(t *testing.T){
    
    testParam := uniq_types.Options{true,false,false,&uniq_types.OptFields{false,0,0}}
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
    want.PushBack(uniq_types.Pair{"I love music.",3})
    want.PushBack(uniq_types.Pair{" ",1})
    want.PushBack(uniq_types.Pair{"I love music of Kartik.",2})

    w := want.Front()
    for e := got.Front(); e != nil; e = e.Next() {
        if e.Value.(uniq_types.Pair).Numb != w.Value.(uniq_types.Pair).Numb {
            require.Equal(t,  e.Value.(uniq_types.Pair).Numb,w.Value.(uniq_types.Pair).Numb, "Must be equal according to the logic of the program")
        }
        if w.Next() != nil {
            w = w.Next()
        }
    }
}
