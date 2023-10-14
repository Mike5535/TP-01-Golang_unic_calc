package uniq

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"strconv"
	"github.com/mike5535/uniq/uniq_types"
	"github.com/mike5535/uniq/cut_string"
)

//функция уникализации строк
func Uniqle(param *uniq_types.Options, input []string) *list.List {
	listPair := new(list.List)
	iterPair := listPair.Front()

	for _,iterInput := range input {

		if iterPair == nil {
			listPair.PushBack(uniq_types.Pair{iterInput, 0})
			iterPair = listPair.Back()
		}

		if curr := iterPair.Value.(uniq_types.Pair); cut_string.CutString(curr.Str, *param.FieldOptions) == cut_string.CutString(iterInput, *param.FieldOptions) {
			curr.Numb++
			iterPair.Value = curr
		} else {
			listPair.PushBack(uniq_types.Pair{iterInput, 1})
			iterPair = iterPair.Next()
		}

	}

	return listPair
}

func UniqWrite(param *uniq_types.Options,listPair *list.List) {
	outName := flag.Arg(1)
	var outFile *os.File
	if (outName != "") {
		f, err := os.OpenFile(outName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		outFile = f
	}

	helpFuncWrite := func (output string) {
		if (outName != "") {
			fmt.Println(outFile.Write([]byte(output+"\n"))) 
		} else {
			fmt.Println(output)
		}
	}

	helpFuncCount := func(e list.Element, flagCounting bool) {
		if (flagCounting) {
			helpFuncWrite(strconv.Itoa(int(e.Value.(uniq_types.Pair).Numb)) + " " + e.Value.(uniq_types.Pair).Str)
		} else {
			helpFuncWrite(e.Value.(uniq_types.Pair).Str)
		}
	}

	for e := listPair.Front(); e != nil; e = e.Next() {
		if param.Repeat && e.Value.(uniq_types.Pair).Numb > 1 {
			helpFuncCount(*e, param.Counting)
		}
		if param.Uniq && e.Value.(uniq_types.Pair).Numb == 1 {
			helpFuncCount(*e, param.Counting)
		}
		if !param.Uniq && !param.Repeat {
			helpFuncCount(*e, param.Counting)
		}
	}

	if (outName != "") {
		outFile.Close()
	}
}
