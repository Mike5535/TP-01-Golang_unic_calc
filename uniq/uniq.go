package uniq

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type OptFields struct {
	AnyRegister bool
	NumFields   int
	NumChars    int
}

type Options struct {
	Counting     bool
	Repeat       bool
	Uniq         bool
	FieldOptions *OptFields
}

func UniqleRead() {
	opt := new(Options)
	opt.FieldOptions = new(OptFields)

	flag.BoolVar(&opt.Counting, "c", false, "strings with number of repeat")
	flag.BoolVar(&opt.Repeat, "d", false, "repeating strings")
	flag.BoolVar(&opt.Uniq, "u", false, "uniq strings")
	flag.BoolVar(&opt.FieldOptions.AnyRegister, "i", false, "ignore register")
	flag.IntVar(&opt.FieldOptions.NumFields, "f", 0, "number of fields to ignore")
	flag.IntVar(&opt.FieldOptions.NumChars, "s", 0, "number of chars to ignore")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)

	orderList := list.New()

	for buf.Scan() {
		orderList.PushBack(buf.Text())
	}

	Uniqle(*opt, *orderList)

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
}

//функция уникализации строк
func Uniqle(param Options, input list.List) {

	type Pair struct {
		str  string
		numb uint
	}

	listPair := new(list.List)
	iterPair := listPair.Front()

	for iterInput := input.Front(); iterInput != nil; iterInput = iterInput.Next() {

		if iterPair == nil {
			listPair.PushBack(Pair{iterInput.Value.(string), 0})
			iterPair = listPair.Back()
		}

		if curr := iterPair.Value.(Pair); cutString(curr.str, *param.FieldOptions) == cutString(iterInput.Value.(string), *param.FieldOptions) {
			curr.numb++
			iterPair.Value = curr
		} else {
			listPair.PushBack(Pair{iterInput.Value.(string), 1})
			iterPair = iterPair.Next()
		}

	}

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
			helpFuncWrite(strconv.Itoa(int(e.Value.(Pair).numb)) + " " + e.Value.(Pair).str)
		} else {
			helpFuncWrite(e.Value.(Pair).str)
		}
	}

	for e := listPair.Front(); e != nil; e = e.Next() {
		if param.Repeat && e.Value.(Pair).numb > 1 {
			helpFuncCount(*e, param.Counting)
		}
		if param.Uniq && e.Value.(Pair).numb == 1 {
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

//убирает символы из строки
func cutSymbols(inputStr string, numChars int) string {
	buf := strings.Split(inputStr, "")
	if buf[0] == " " {
		copy(buf[:], buf[1:])
		buf[len(buf)-1] = ""
		buf = buf[:len(buf)-1]
	}
	buf = buf[numChars:]
	return strings.Join(buf, "")
}

//убирает поля из строки
func cutString(inputStr string, param OptFields) string {
	if inputStr == "" {
		return ""
	}

	if param.AnyRegister {
		inputStr = strings.ToLower(inputStr)
	}

	if param.NumFields != 0 {
		s := strings.Split(inputStr, " ")[param.NumFields:]
		temp := strings.Join(s, "")
		if param.NumChars != 0 {
			temp = cutSymbols(temp, param.NumChars)
		}
		return temp
	}

	if param.NumChars != 0 {
		return cutSymbols(inputStr, param.NumChars)
	}

	return inputStr
}
