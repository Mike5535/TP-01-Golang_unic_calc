package uniq_read

import (
	"fmt"
	"bufio"
	"container/list"
	"flag"
	"io"
	"os"
	"github.com/mike5535/uniq/uniq_types"
)

func UniqleRead() (*uniq_types.Options, *list.List) {
	opt := new(uniq_types.Options)
	opt.FieldOptions = new(uniq_types.OptFields)

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

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
	return opt, orderList
}
